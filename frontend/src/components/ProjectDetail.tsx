import React, {
  useState,
  useCallback,
  useEffect,
  memo,
  lazy,
  Suspense,
} from "react";
import { DraggableModal } from "../pkg/DraggableModal";
import { detailProject } from "../api/project";
import { Button, Tabs, Skeleton, Badge } from "antd";
import DeployStatus from "./DeployStatus";
import { setNamespaceReload } from "../store/actions";
import ErrorBoundary from "./ErrorBoundary";
import ServiceEndpoint from "./ServiceEndpoint";
import { useDispatch } from "react-redux";
import pb from "../api/compiled";

import TabLog from "./TabLog";
const TabInfo = lazy(() => import("./TabInfo"));
const TabEdit = lazy(() => import("./TabEdit"));
const Shell = lazy(() => import("./TabShell"));

const { TabPane } = Tabs;

const ItemDetailModal: React.FC<{
  item: pb.types.ProjectModel;
  namespace: string;
  namespaceId: number;
}> = ({ item, namespace, namespaceId }) => {
  const dispatch = useDispatch();
  const [visible, setVisible] = useState(false);
  const onOk = useCallback(() => setVisible(true), []);
  const [detail, setDetail] = useState<pb.project.ShowResponse | undefined>();
  const [resizeAt, setResizeAt] = useState<number>(0);

  useEffect(() => {
    if (visible && namespaceId && item.id) {
      detailProject(item.id).then((res) => {
        setDetail(res.data);
      });
    }
  }, [item.id, visible, namespaceId]);

  const onSuccess = useCallback(() => {
    item.id && detailProject(item.id).then((res) => {
      setDetail(res.data);
    });
  }, [item.id]);

  const onCancel = useCallback(() => {
    setVisible(false);
  }, []);

  return (
    <div className="project-detail">
      <Button
        onClick={() => {
          onOk();
        }}
        className="project-detail__show-button"
        type="dashed"
      >
        <DeployStatus status={item.deploy_status} />
        <span
          title={item.name}
          style={{
            textOverflow: "ellipsis",
            whiteSpace: "nowrap",
            overflow: "hidden",
            marginRight: 5,
          }}
        >
          {item.name}
        </span>
        {item.deploy_status === pb.types.Deploy.StatusDeployed && (
          <ServiceEndpoint projectId={item.id} />
        )}
      </Button>
      <DraggableModal
        onResize={() => {
          setResizeAt(new Date().getTime());
        }}
        className="draggable-modal"
        visible={visible}
        initialWidth={900}
        initialHeight={600}
        footer={null}
        keyboard={false}
        onCancel={onCancel}
        title={
          <Badge.Ribbon
            className="project-detail__badge"
            placement="start"
            text={namespace}
          >
            <div style={{ textAlign: "center", fontSize: 18 }}>{item.name}</div>
          </Badge.Ribbon>
        }
      >
        <Tabs
          destroyInactiveTabPane
          defaultActiveKey="1"
          centered
          style={{ height: "100%" }}
        >
          {(item.deploy_status === pb.types.Deploy.StatusDeployed || item.deploy_status === pb.types.Deploy.StatusDeploying) && (
            <>
              <TabPane tab="容器日志" key="container-logs">
                {detail?.project && detail.project.namespace ? (
                  <TabLog
                    updatedAt={detail.project.updated_at}
                    id={detail.project.id}
                    namespace={detail.project.namespace.name}
                  />
                ) : (
                  <Skeleton active />
                )}
              </TabPane>
              <TabPane tab="命令行" key="shell" style={{ height: "100%" }}>
                <Suspense fallback={<Skeleton active />}>
                  <ErrorBoundary>
                    {detail?.project && detail.project.namespace && (
                      <Shell
                        namespace={detail.project.namespace.name}
                        id={detail.project.id}
                        updatedAt={detail.project.updated_at}
                        resizeAt={resizeAt}
                      />
                    )}
                  </ErrorBoundary>
                </Suspense>
              </TabPane>
              <TabPane tab="配置更新" key="update-config">
                <Suspense fallback={<Skeleton active />}>
                  {detail?.project && detail.project.namespace && (
                    <TabEdit
                      elements={detail.elements}
                      namespaceId={detail.project.namespace.id}
                      detail={detail.project}
                      updatedAt={detail.project.updated_at}
                      onSuccess={onSuccess}
                    />
                  )}
                </Suspense>
              </TabPane>
            </>
          )}
          <TabPane tab="详细信息" key="detail" className="detail-tab">
            <Suspense fallback={<Skeleton active />}>
              {detail?.project && (
                <TabInfo
                  detail={detail.project}
                  cpu={detail.cpu}
                  memory={detail.memory}
                  git_commit_web_url={detail.project.git_commit_web_url}
                  git_commit_title={detail.project.git_commit_title}
                  git_commit_author={detail.project.git_commit_author}
                  git_commit_date={detail.project.git_commit_date}
                  urls={detail.urls}
                  onDeleted={() => {
                    dispatch(setNamespaceReload(true));
                    setVisible(false);
                  }}
                />
              )}
            </Suspense>
          </TabPane>
        </Tabs>
      </DraggableModal>
    </div>
  );
};

export default memo(ItemDetailModal);
