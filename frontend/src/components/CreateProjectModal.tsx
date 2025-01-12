import React, { useState, useCallback, useEffect, useMemo, memo } from "react";
import { selectClusterInfo } from "../store/reducers/cluster";
import PipelineInfo from "./PipelineInfo";
import Elements from "./elements/Elements";
import { DraggableModal } from "../pkg/DraggableModal/DraggableModal";
import { MyCodeMirror as CodeMirror, getMode } from "./MyCodeMirror";
import pb from "../api/compiled";
import { orderBy } from "lodash";
import { useAsyncState } from "../utils/async";

import { configFile } from "../api/git";
import {
  DeployStatus as DeployStatusEnum,
  selectList,
} from "../store/reducers/createProject";
import { useWs, useWsReady } from "../contexts/useWebsocket";
import { message, Progress, Button, Form } from "antd";
import {
  PlusOutlined,
  StopOutlined,
  ArrowLeftOutlined,
  ArrowRightOutlined,
} from "@ant-design/icons";
import { useDispatch, useSelector } from "react-redux";
import {
  clearCreateProjectLog,
  setCreateProjectLoading,
  setDeployStatus,
} from "../store/actions";
import classNames from "classnames";
import { toSlug } from "../utils/slug";
import LogOutput from "./LogOutput";
import ProjectSelector from "./ProjectSelector";
import DebugModeSwitch from "./DebugModeSwitch";
import TimeCost from "./TimeCost";

const initFormValues = {
  debug: true,
  extra_values: [],
  config: "",
};

const CreateProjectModal: React.FC<{
  namespaceId: number;
}> = ({ namespaceId }) => {
  const list = useSelector(selectList);
  const dispatch = useDispatch();
  const [form] = Form.useForm();
  const [mode, setMode] = useState<string>("text/x-yaml");
  const [visible, setVisible] = useAsyncState<boolean>(false);
  const [editVisible, setEditVisible] = useState<boolean>(true);
  const [timelineVisible, setTimelineVisible] = useState<boolean>(false);
  const ws = useWs();
  const wsReady = useWsReady();
  const [start, setStart] = useState(false);
  const info = useSelector(selectClusterInfo);
  const [data, setData] = useState<{
    projectName: string;
    gitProjectId: number;
    gitBranch: string;
    gitCommit: string;
  }>();
  const [elements, setElements] = useState<pb.mars.Element[]>();

  let slug = useMemo(
    () => toSlug(namespaceId, data?.projectName ? data.projectName : ""),
    [namespaceId, data]
  );

  const onCancel = useCallback(() => {
    setElements([]);
    form.resetFields();
    setData(undefined);
    setVisible(false);
    setEditVisible(true);
    setTimelineVisible(false);
    dispatch(clearCreateProjectLog(slug));
  }, [dispatch, slug, form, setVisible]);

  const onOk = useCallback(
    (values: any) => {
      if (values.extra_values) {
        let defaults = elements
          ? elements.map((i) => ({
              path: i.path,
              value: i.default,
            }))
          : [];
        values.extra_values =
          values.extra_values.length > 0
            ? values.extra_values.map((i: any) => ({
                ...i,
                value: String(i.value),
              }))
            : defaults;
      }
      if (!wsReady) {
        message.error("连接断开了");
        return;
      }
      if (data && data.gitProjectId && data.gitBranch && data.gitCommit) {
        // todo ws connected!
        setEditVisible(false);
        setTimelineVisible(true);
        let s = pb.websocket.CreateProjectInput.encode({
          type: pb.websocket.Type.CreateProject,
          namespace_id: Number(namespaceId),
          name: data.projectName,
          git_project_id: Number(data.gitProjectId),
          git_branch: data.gitBranch,
          git_commit: data.gitCommit,
          config: values.config,
          atomic: !values.debug,
          extra_values: values.extra_values,
        }).finish();

        dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));

        dispatch(clearCreateProjectLog(slug));
        dispatch(setCreateProjectLoading(slug, true));
        setStart(true);
        ws?.send(s);
        return;
      }

      message.error("项目id, 分支，提交必填");
    },
    [dispatch, slug, ws, namespaceId, wsReady, data, elements]
  );

  const onRemove = useCallback(() => {
    if (data && data.gitProjectId && data.gitBranch && data.gitCommit) {
      let s = pb.websocket.CancelInput.encode({
        type: pb.websocket.Type.CancelProject,
        namespace_id: namespaceId,
        name: data.projectName,
      }).finish();
      ws?.send(s);
      return;
    }
  }, [ws, namespaceId, data]);

  const loadConfigFile = useCallback(
    (gitProjectId: string, gitBranch: string) => {
      configFile({
        git_project_id: gitProjectId,
        branch: gitBranch,
      }).then((res) => {
        if (!form.getFieldValue("config")) {
          form.setFieldsValue({ config: res.data.data });
        }
        setMode(getMode(res.data.type));
        setElements(res.data.elements);
      });
    },
    [form]
  );

  const onChange = useCallback(
    (v: any) => {
      setData(v);
      form.setFieldsValue({ selectors: v });
      if (v.gitCommit !== "") {
        loadConfigFile(String(v.gitProjectId), v.gitBranch);
      }
    },
    [form, loadConfigFile]
  );

  useEffect(() => {
    if (list[slug]?.deployStatus !== DeployStatusEnum.DeployUnknown) {
      setStart(false);
    }
    if (list[slug]?.deployStatus === DeployStatusEnum.DeploySuccess) {
      setTimelineVisible(false);
      setEditVisible(true);
      setElements([]);
      form.resetFields();
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));
      setTimeout(() => {
        setVisible(false);
      }, 500);
    }
  }, [list, dispatch, slug, onCancel, form, setVisible]);

  useEffect(() => {
    if (!wsReady) {
      setStart(false);
      dispatch(setCreateProjectLoading(slug, false));
    }
  }, [wsReady, dispatch, slug]);

  return (
    <div>
      <Button
        onClick={() => setVisible(true)}
        style={{ width: "100%" }}
        type="dashed"
        icon={<PlusOutlined />}
      ></Button>
      <DraggableModal
        destroyOnClose
        visible={visible}
        okButtonProps={{
          loading: list[slug]?.isLoading,
          danger: info.status === "bad",
        }}
        onCancel={onCancel}
        cancelButtonProps={{ disabled: list[slug]?.isLoading }}
        closable={!list[slug]?.isLoading}
        footer={null}
        initialWidth={900}
        initialHeight={600}
        title={<div style={{ textAlign: "center" }}>创建项目</div>}
        className="draggable-modal drag-item-modal"
      >
        <div
          className="create-project-modal"
          style={{ display: "flex", flexDirection: "column", height: "100%" }}
        >
          {data?.gitCommit && (
            <PipelineInfo
              projectId={data.gitProjectId}
              branch={data.gitBranch}
              commit={data.gitCommit}
            />
          )}
          <div
            style={{ height: "100%" }}
            className={classNames({ "display-none": !editVisible })}
          >
            <Form
              layout="horizontal"
              form={form}
              labelWrap
              autoComplete="off"
              onFinish={onOk}
              initialValues={initFormValues}
              style={{
                height: "100%",
                display: "flex",
                flexDirection: "column",
              }}
            >
              <div
                style={{
                  display: "flex",
                  alignItems: "center",
                  marginBottom: 10,
                }}
              >
                {list[slug]?.output?.length > 0 && (
                  <Button
                    style={{ marginRight: 5 }}
                    type="dashed"
                    disabled={list[slug]?.isLoading}
                    onClick={() => {
                      setEditVisible(false);
                      setTimelineVisible(true);
                    }}
                    icon={<ArrowRightOutlined />}
                  />
                )}
                <Form.Item
                  name="selectors"
                  style={{ width: "100%", margin: 0 }}
                  rules={[{ required: true, message: "项目必选" }]}
                >
                  <ProjectSelector isCreate onChange={onChange} />
                </Form.Item>
              </div>

              <div
                style={{
                  display: "flex",
                  justifyContent: "space-between",
                  alignItems: "center",
                  paddingBottom: 10,
                }}
              >
                <div style={{ display: "flex" }}>
                  <span style={{ marginRight: 5 }}>配置文件:</span>

                  <Button
                    htmlType="submit"
                    style={{ fontSize: 12, marginRight: 5 }}
                    size="small"
                    danger={info.status === "bad"}
                    type={"primary"}
                    loading={list[slug]?.isLoading}
                  >
                    {info.status === "bad" ? "集群资源不足" : "部署"}
                  </Button>
                </div>
                <Form.Item noStyle name={"debug"}>
                  <DebugModeSwitch />
                </Form.Item>
              </div>
              <div
                style={{
                  minWidth: 200,
                  marginBottom: 20,
                  height: "100%",
                }}
              >
                <Form.Item name="extra_values" noStyle>
                  <Elements
                    elements={orderBy(elements, ["type"], ["asc"])}
                    style={{
                      inputNumber: { fontSize: 10, width: "100%" },
                      input: { fontSize: 10 },
                      label: { fontSize: 10 },
                      formItem: {
                        marginBottom: 5,
                        display: "inline-block",
                        width: "calc(33% - 8px)",
                        marginRight: 8,
                      },
                    }}
                  />
                </Form.Item>
                <Form.Item name="config" style={{ height: "100%" }} noStyle>
                  <CodeMirror
                    options={{
                      mode: mode,
                      theme: "dracula",
                      lineNumbers: true,
                    }}
                  />
                </Form.Item>
              </div>
            </Form>
          </div>

          <div
            id="preview"
            style={{ height: "100%", overflow: "auto" }}
            className={classNames("preview", {
              "display-none": !timelineVisible,
            })}
          >
            <div
              style={{
                display: "flex",
                alignItems: "center",
                marginBottom: 20,
              }}
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
                strokeColor={{
                  from: "#108ee9",
                  to: "#87d068",
                }}
                style={{ padding: "0 10px" }}
                percent={list[slug]?.processPercent}
                status="active"
              />
            </div>

            <div
              style={{
                display: "flex",
                alignItems: "center",
                marginBottom: 10,
              }}
            >
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
              <TimeCost start={start} />
            </div>

            <LogOutput slug={slug} />
          </div>
        </div>
      </DraggableModal>
    </div>
  );
};

export default memo(CreateProjectModal);
