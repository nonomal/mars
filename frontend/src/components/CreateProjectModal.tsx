import React, { useState, useRef, useCallback, useEffect } from "react";
import PipelineInfo from "./PipelineInfo";
import { DraggableModal } from "ant-design-draggable-modal";
import { Controlled as CodeMirror } from "react-codemirror2";
import {
  branches,
  commits,
  configFile,
  Options,
  projects,
} from "../api/gitlab";
import {
  DeployStatus as DeployStatusEnum,
  selectList,
} from "../store/reducers/createProject";
import _ from "lodash";
import { CascaderOptionType } from "antd/lib/cascader";
import { useWs } from "../contexts/useWebsocket";
import { message, Progress } from "antd";
import { Button, Cascader, Timeline } from "antd";
import {
  PlusOutlined,
  ArrowLeftOutlined,
  ArrowRightOutlined,
} from "@ant-design/icons";
import "codemirror/lib/codemirror.css";
import "codemirror/theme/material.css";
import "codemirror/theme/dracula.css";
import { useDispatch, useSelector } from "react-redux";
import {
  clearCreateProjectLog,
  setCreateProjectLoading,
  setDeployStatus,
} from "../store/actions";
import classNames from "classnames";
import { toSlug } from "../utils/slug";

require("codemirror/mode/go/go");
require("codemirror/mode/css/css");
require("codemirror/mode/javascript/javascript");
require("codemirror/mode/yaml/yaml");
require("codemirror/mode/php/php");
require("codemirror/mode/textile/textile");

const initItemData: CreateItemInterface = {
  name: "",
  gitlabProjectId: 0,
  gitlabBranch: "",
  gitlabCommit: "",
  config: "",
};

interface CreateItemInterface {
  gitlabProjectId: number;
  gitlabBranch: string;
  gitlabCommit: string;

  name: string;
  config: string;
}

const CreateProjectModal: React.FC<{
  namespaceId: number;
}> = ({ namespaceId }) => {
  const list = useSelector(selectList);
  const dispatch = useDispatch();
  const [data, setData] = useState<CreateItemInterface>(initItemData);
  const [mode, setMode] = useState<string>("text/x-yaml");
  const [options, setOptions] = useState<Options[]>([]);
  const [visible, setVisible] = useState<boolean>(false);
  const [editVisible, setEditVisible] = useState<boolean>(true);
  const [timelineVisible, setTimelineVisible] = useState<boolean>(false);
  const [value, setValue] = useState<string[]>([]);

  let slug = toSlug(namespaceId, data.name);

  const onCancel = useCallback(() => {
    setValue([]);
    setVisible(false);
    setEditVisible(true);
    setTimelineVisible(false);
    setData(initItemData);
    dispatch(clearCreateProjectLog(slug));
  }, [dispatch, slug]);

  useEffect(() => {
    projects().then((res) => setOptions(res.data.data));
  }, []);

  useEffect(() => {
    if (list[slug]?.deployStatus === DeployStatusEnum.DeploySuccess) {
      setTimelineVisible(false);
      setEditVisible(true);
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));
      setTimeout(() => {
        setVisible(false);
        setValue([]);
        setData(initItemData);
      }, 500);
    }
  }, [list, dispatch, slug]);

  const onChange = (
    values: any[],
    selectedOptions: CascaderOptionType[] | undefined
  ) => {
    let gitlabId = _.get(values, 0, 0);
    let gbranch = _.get(values, 1, "");
    let gcommit = _.get(values, 2, "");
    setData((d) => ({
      ...d,
      name: _.get(
        options.find((item) => item.value === values[0]),
        "label",
        ""
      ),
      gitlabProjectId: gitlabId,
      gitlabBranch: gbranch,
      gitlabCommit: gcommit,
    }));

    if (selectedOptions) {
      const targetOption = selectedOptions[selectedOptions.length - 1];
      if (targetOption.children) {
        targetOption.loading = true;
        targetOption.children = undefined;
        switch (targetOption.type) {
          case "project":
            branches(Number(targetOption.value)).then((res) => {
              targetOption.loading = false;
              targetOption.children = res.data.data;
              setOptions([...options]);
            });
            return;
          case "branch":
            commits(
              Number(targetOption.projectId),
              String(targetOption.value)
            ).then((res) => {
              targetOption.loading = false;
              targetOption.children = res.data.data;
              setOptions([...options]);
            });
            return;
        }
      }
    }

    if (gitlabId) {
      let o = options.find((item) => item.value === values[0]);
      setValue([o ? o.label : ""]);
      if (gbranch) {
        if (o && o.children) {
          let b = o.children.find((item) => item.value === gbranch);
          setValue([o.label, b ? b.label : ""]);
          if (gcommit) {
            if (b && b.children) {
              let c = b.children.find((item) => item.value === gcommit);
              setValue([o.label, b.label, c ? c.label : ""]);
              if (data.config === "") {
                loadConfigFile();
              }
            }
          }
        }
      }
    }
  };
  const cmref = useRef<any>();

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
  }, [data.gitlabBranch, data.gitlabProjectId]);

  useEffect(() => {
    if (cmref.current && data.config) {
      cmref.current.editor.setSize("100%", "100%");
    }
  }, [data.config]);

  const loadData = (selectedOptions: CascaderOptionType[] | undefined) => {
    if (!selectedOptions) {
      return;
    }
    const targetOption = selectedOptions[selectedOptions.length - 1];
    targetOption.loading = true;

    console.log(targetOption);

    switch (targetOption.type) {
      case "project":
        branches(Number(targetOption.value)).then((res) => {
          targetOption.loading = false;
          targetOption.children = res.data.data;
          setOptions([...options]);
        });
        return;
      case "branch":
        commits(
          Number(targetOption.projectId),
          String(targetOption.value)
        ).then((res) => {
          targetOption.loading = false;
          targetOption.children = res.data.data;
          setOptions([...options]);
        });
        return;
    }
  };
  const ws = useWs();

  const onOk = useCallback(() => {
    console.log(data);
    if (data.gitlabProjectId && data.gitlabBranch && data.gitlabCommit) {
      // todo ws connected!
      setEditVisible(false);
      setTimelineVisible(true);

      let re = {
        type: "create_project",
        data: JSON.stringify({
          namespace_id: Number(namespaceId),
          name: data.name,
          gitlab_project_id: Number(data.gitlabProjectId),
          gitlab_branch: data.gitlabBranch,
          gitlab_commit: data.gitlabCommit,
          config: data.config,
        }),
      };

      let s = JSON.stringify(re);
      dispatch(clearCreateProjectLog(slug));
      dispatch(setCreateProjectLoading(slug, true));
      ws?.send(s);
      return;
    }

    message.error("项目id, 分支，提交必填");
  }, [data, dispatch, slug, ws, namespaceId]);

  return (
    <div>
      <Button
        onClick={() => setVisible(true)}
        style={{ width: "100%" }}
        type="dashed"
        icon={<PlusOutlined />}
      ></Button>
      <DraggableModal
        visible={visible}
        okButtonProps={{ loading: list[slug]?.isLoading }}
        cancelButtonProps={{ disabled: list[slug]?.isLoading }}
        closable={!list[slug]?.isLoading}
        okText="部署"
        cancelText="取消"
        onOk={onOk}
        initialWidth={800}
        initialHeight={500}
        title="创建项目"
        className="drag-item-modal"
        onCancel={onCancel}
      >
        <PipelineInfo
          projectId={data.gitlabProjectId}
          branch={data.gitlabBranch}
          commit={data.gitlabCommit}
        />
        <div className={classNames({ "display-none": !editVisible })}>
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
          <Cascader
            options={options}
            style={{ width: "100%", marginBottom: "10px" }}
            autoFocus
            value={value}
            allowClear={false}
            loadData={loadData}
            onChange={onChange}
            changeOnSelect
            placeholder="选择项目/分支/提交"
          />
          配置文件:
          <div
            style={{
              minWidth: 200,
              maxWidth: 1280,
              marginBottom: 20,
              height: "100%",
            }}
          >
            <CodeMirror
              ref={cmref}
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

          <Timeline
            pending={list[slug]?.isLoading ? "loading..." : false}
            reverse={true}
            style={{ paddingLeft: 2 }}
          >
            {list[slug]?.output.map((data, index) => (
              <Timeline.Item
                key={index}
                color={data === "部署失败" ? "red" : "blue"}
              >
                {data}
              </Timeline.Item>
            ))}
          </Timeline>
        </div>
      </DraggableModal>
    </div>
  );
};

export default CreateProjectModal;
