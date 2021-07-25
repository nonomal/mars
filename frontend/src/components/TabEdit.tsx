import React, { useState, useEffect, useCallback, memo } from "react";
import { Controlled as CodeMirror } from "react-codemirror2";
import PipelineInfo from "./PipelineInfo";

import { commit, configFile, projects } from "../api/gitlab";
import {
  DeployStatus as DeployStatusEnum,
  selectList,
} from "../store/reducers/createProject";
import { Button, Skeleton, Progress } from "antd";
import "codemirror/lib/codemirror.css";
import "codemirror/theme/material.css";
import "codemirror/theme/dracula.css";
import { useSelector, useDispatch } from "react-redux";
import {
  clearCreateProjectLog,
  resetElapsedTime,
  setCreateProjectLoading,
  setDeployStatus,
} from "../store/actions";
import { toSlug } from "../utils/slug";
import { useWs } from "../contexts/useWebsocket";
import {
  ArrowLeftOutlined,
  StopOutlined,
  ArrowRightOutlined,
} from "@ant-design/icons";
import classNames from "classnames";
import { ProjectDetail } from "../api/project";
import LogOutput from "./LogOutput";
import ProjectSelector from "./ProjectSelector";
import TimeCost from "./TimeCost";

require("codemirror/mode/go/go");
require("codemirror/mode/css/css");
require("codemirror/mode/javascript/javascript");
require("codemirror/mode/yaml/yaml");
require("codemirror/mode/php/php");
require("codemirror/mode/textile/textile");

interface CreateItemInterface {
  gitlabProjectId: number;
  gitlabBranch: string;
  gitlabCommit: string;

  name: string;
  config: string;
}

const ModalSub: React.FC<{
  detail: ProjectDetail;
  onSuccess: () => void;
}> = ({ detail, onSuccess }) => {
  let id = detail.id;
  let namespaceId = detail.namespace.id;
  const ws = useWs();

  const [editVisible, setEditVisible] = useState<boolean>(true);
  const [timelineVisible, setTimelineVisible] = useState<boolean>(false);
  const list = useSelector(selectList);
  const dispatch = useDispatch();
  const [data, setData] = useState<CreateItemInterface>({
    name: detail.name,
    gitlabProjectId: Number(detail.gitlab_project_id),
    gitlabBranch: detail.gitlab_branch,
    gitlabCommit: detail.gitlab_commit,
    config: detail.config,
  });
  const [mode, setMode] = useState<string>("text/x-yaml");
  const [initValue, setInitValue] = useState<{
    projectName: string;
    gitlabProjectId: number;
    gitlabBranch: string;
    gitlabCommit: string;
    time?: number;
  }>();
  let slug = toSlug(namespaceId, data.name);

  // 初始化，设置 initvalue
  useEffect(() => {
    projects().then((res) => {
      if (
        detail &&
        detail.gitlab_project_id &&
        detail.gitlab_branch &&
        detail.gitlab_commit
      ) {
        commit(
          Number(detail.gitlab_project_id),
          detail.gitlab_branch,
          detail.gitlab_commit
        ).then((res) => {
          setInitValue({
            projectName: detail.name,
            gitlabProjectId: Number(detail.gitlab_project_id),
            gitlabBranch: detail.gitlab_branch,
            gitlabCommit: res.data.data.label,
          });
        });
      }
    });
  }, [setInitValue, detail]);

  // 更新成功，触发 onSuccess
  useEffect(() => {
    if (list[slug]?.deployStatus === DeployStatusEnum.DeployUpdateSuccess) {
      setTimelineVisible(false);
      setEditVisible(true);
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));
      onSuccess();
    }
  }, [list, dispatch, slug, onSuccess]);

  // 更新 config 文件的类型， TODO 支持动态加载 mode css 文件
  const loadConfigFile = useCallback(() => {
    configFile(data.gitlabProjectId, data.gitlabBranch).then((res) => {
      setData((d) => ({ ...d, config: res.data.data.data }));
      switch (res.data.data.type) {
        case "dotenv":
        case "env":
        case ".env":
          setMode("text/x-textile");
          break;
        case "yaml":
          setMode("text/x-yaml");
          break;
        case "php":
          setMode("php");
          break;
        default:
          setMode(res.data.data.type);
          break;
      }
    });
  }, [data.gitlabProjectId, data.gitlabBranch]);

  const onChange = ({
    projectName,
    gitlabProjectId,
    gitlabBranch,
    gitlabCommit,
  }: {
    projectName: string;
    gitlabProjectId: number;
    gitlabBranch: string;
    gitlabCommit: string;
  }) => {
    setData((d) => ({
      ...d,
      name: projectName,
      gitlabProjectId: gitlabProjectId,
      gitlabBranch: gitlabBranch,
      gitlabCommit: gitlabCommit,
    }));

    if (gitlabCommit !== "" && data.config === "") {
      loadConfigFile();
    }
  };
  const updateDeploy = () => {
    if (data.gitlabCommit && data.gitlabBranch) {
      setEditVisible(false);
      setTimelineVisible(true);

      let re = {
        type: "update_project",
        data: JSON.stringify({
          project_id: Number(id),
          gitlab_branch: data.gitlabBranch,
          gitlab_commit: data.gitlabCommit,
          config: data.config,
        }),
      };
      let s = JSON.stringify(re);
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));
      dispatch(resetElapsedTime(slug));

      dispatch(clearCreateProjectLog(slug));
      dispatch(setCreateProjectLoading(slug, true));
      ws?.send(s);
    }
  };

  const onReset = () => {
    setData({
      name: detail.name,
      gitlabProjectId: Number(detail.gitlab_project_id),
      gitlabBranch: detail.gitlab_branch,
      gitlabCommit: detail.gitlab_commit,
      config: detail.config,
    });
    if (initValue) {
      setInitValue({ ...initValue, time: new Date().getUTCSeconds() });
    }
  };

  const onRemove = useCallback(() => {
    if (data.gitlabProjectId && data.gitlabBranch && data.gitlabCommit) {
      let re = {
        type: "cancel_project",
        data: JSON.stringify({
          namespace_id: Number(namespaceId),
          name: data.name,
        }),
      };

      let s = JSON.stringify(re);
      ws?.send(s);
      return;
    }
  }, [data, ws, namespaceId]);

  return (
    <div className="edit-project">
      <PipelineInfo
        projectId={data.gitlabProjectId}
        branch={data.gitlabBranch}
        commit={data.gitlabCommit}
      />
      <div className={classNames({ "display-none": !editVisible })}>
        <div style={{ width: "100%" }}>
          {list[slug]?.output.length > 0 ? (
            <Button
              style={{ marginBottom: 20 }}
              type="dashed"
              disabled={list[slug]?.isLoading}
              onClick={() => {
                setEditVisible(false);
                setTimelineVisible(true);
              }}
              icon={<ArrowRightOutlined />}
            />
          ) : (
            ""
          )}
          {initValue ? (
            <ProjectSelector value={initValue} onChange={onChange} />
          ) : (
            <Skeleton.Input active style={{ width: 800 }} size="small" />
          )}
        </div>
        配置文件:
        <div style={{ minWidth: 200, marginBottom: 20 }}>
          <CodeMirror
            value={data.config}
            options={{
              mode: mode,
              theme: "dracula",
              lineNumbers: true,
            }}
            onBeforeChange={(editor, d, value) => {
              console.log(editor, d, value);
              setData({ ...data, config: value });
            }}
          />
        </div>
      </div>
      <div
        id="preview"
        style={{ height: "100%", overflow: "auto" }}
        className={classNames("preview", {
          "display-none": !timelineVisible,
        })}
      >
        <div
          style={{ display: "flex", alignItems: "center", marginBottom: 20 }}
        >
          <Button
            type="dashed"
            disabled={list[slug]?.isLoading}
            onClick={() => {
              setEditVisible(true);
              setTimelineVisible(false);
            }}
            icon={<ArrowLeftOutlined />}
          />
          <Progress
            style={{ padding: "0 10px" }}
            percent={list[slug]?.processPercent}
            status="active"
          />
        </div>
        <div style={{display: "flex", alignItems: "center", marginBottom: 10}}>
          <Button
            hidden={
              list[slug]?.deployStatus === DeployStatusEnum.DeployCanceled
            }
            style={{ marginRight: 10 }}
            danger
            icon={<StopOutlined />}
            type="dashed"
            onClick={onRemove}
          >
            取消
          </Button>
          <TimeCost seconds={list[slug]?.ElapsedTime} />
        </div>
        <LogOutput slug={slug} />
      </div>

      <div className="edit-project__footer">
        <Button
          style={{ marginRight: 5 }}
          disabled={list[slug]?.isLoading}
          onClick={onReset}
        >
          重置
        </Button>
        <Button
          type="primary"
          loading={list[slug]?.isLoading}
          onClick={updateDeploy}
        >
          部署
        </Button>
      </div>
    </div>
  );
};

export default memo(ModalSub);
