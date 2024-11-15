import React, { useState, memo, useCallback } from "react";
import {
  Card,
  Popconfirm,
  Row,
  Spin,
  Col,
  message,
  Tooltip,
  Button,
  Space,
  Popover,
  Form,
  Input,
  Modal,
  Alert,
} from "antd";
import yaml from "js-yaml";
import "../pkg/DraggableModal/index.css";
import {
  CloseOutlined,
  EditFilled,
  LockOutlined,
  UnlockOutlined,
} from "@ant-design/icons";
import ServiceEndpoint from "./ServiceEndpoint";
import ProjectDetail from "./ProjectDetail";
import CreateProjectModal from "./CreateProjectModal";
import { copy } from "../utils/copy";
import styled from "@emotion/styled";
import { useAuth } from "../contexts/auth";
import { components } from "../api/schema";
import ajax from "../api/ajax";
import IconFont from "./Icon";
import TextArea from "antd/es/input/TextArea";
import { css } from "@emotion/css";
import { MyCodeMirror } from "./MyCodeMirror";
import DiffViewer from "./DiffViewer";
import CpuMemory from "./CpuMemory";

const Item: React.FC<{
  item: components["schemas"]["types.NamespaceModel"];
  onNamespaceDeleted: () => void;
  onFavorite: (nsID: number, favorite: boolean) => void;
  loading: boolean;
  reload: () => void;
}> = ({ item, onNamespaceDeleted, loading, onFavorite, reload }) => {
  const [deleting, setDeleting] = useState<boolean>(false);
  const [editDesc, setEditDesc] = useState(false);
  const [popoverVisible, setPopoverVisible] = useState(false);

  return (
    <Card
      style={{ height: "100%" }}
      title={
        <div style={{ marginTop: 8, marginBottom: 5 }}>
          <Row>
            <Col
              span={18}
              offset={3}
              style={{
                display: "flex",
                justifyContent: "center",
              }}
            >
              <IconFont
                onClick={() => onFavorite(item.id, !item.favorite)}
                name="#icon-wodeguanzhu"
                className={css`
                  margin-right: 3px;
                  margin-top: 3px;
                  transition: all 0.3s ease;
                  &:hover {
                    transform: scale(1.2);
                  }
                `}
                style={{
                  color: !item.favorite ? "gray" : "#a78bfa",
                  cursor: "pointer",
                }}
              />
              <Tooltip
                title={<span style={{ fontSize: 10 }}>id: {item.id}</span>}
              >
                <TitleNamespace onClick={() => copy(item.id, "已复制 id")}>
                  项目空间: <TitleNamespaceName>{item.name}</TitleNamespaceName>
                </TitleNamespace>
              </Tooltip>
            </Col>
            <Col span={3} style={{ textAlign: "right" }}>
              {useIsOwned(item) && (
                <Popconfirm
                  title={`确定要删除 '${item.name}' 这个名称空间吗？`}
                  okText="Yes"
                  cancelText="No"
                  onConfirm={() => {
                    setDeleting(true);
                    ajax
                      .DELETE("/api/namespaces/{id}", {
                        params: { path: { id: item.id } },
                      })
                      .then(({ error }) => {
                        if (error) {
                          message.error(error.message);
                          return;
                        }
                        message.success("删除成功");
                        onNamespaceDeleted();
                      });
                  }}
                >
                  <Button
                    type="link"
                    danger
                    size="small"
                    style={{ border: 0 }}
                    icon={<CloseOutlined />}
                  />
                </Popconfirm>
              )}
            </Col>
          </Row>

          <Row>
            <Col
              offset={8}
              span={8}
              style={{
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
              }}
            >
              <Popover
                destroyTooltipOnHide
                trigger="click"
                onOpenChange={(v) => {
                  setPopoverVisible(v);
                  if (!v) {
                    setEditDesc(false);
                  }
                }}
                content={
                  editDesc ? (
                    <Form
                      name="basic"
                      style={{ width: 300 }}
                      initialValues={{ id: item.id, desc: item.description }}
                      autoComplete="off"
                      onFinish={(values) => {
                        console.log(values);
                        ajax
                          .POST("/api/namespaces/{id}/update_desc", {
                            params: { path: { id: values.id } },
                            body: values,
                          })
                          .then(({ error }) => {
                            if (error) {
                              return;
                            }
                            reload();
                            setEditDesc(false);
                            message.success("修改成功");
                          });
                      }}
                    >
                      <Form.Item<
                        components["schemas"]["namespace.UpdateDescRequest"]
                      >
                        name="id"
                        hidden
                      >
                        <Input />
                      </Form.Item>
                      <Form.Item<
                        components["schemas"]["namespace.UpdateDescRequest"]
                      >
                        name="desc"
                        style={{ marginBottom: 5 }}
                      >
                        <TextArea style={{ fontSize: 10 }} rows={5} />
                      </Form.Item>
                      <div style={{ textAlign: "right" }}>
                        <Button
                          style={{ fontSize: 10 }}
                          size="small"
                          htmlType="submit"
                        >
                          提交
                        </Button>
                      </div>
                    </Form>
                  ) : (
                    <div
                      style={{
                        width: 300,
                        fontSize: 12,
                        textAlign: "right",
                      }}
                    >
                      <div
                        style={{
                          textAlign: "left",
                          whiteSpace: "pre-wrap",
                          fontSize: 12,
                        }}
                      >
                        {item.description}
                      </div>
                      <EditFilled
                        style={{ textAlign: "right" }}
                        onClick={() => setEditDesc(true)}
                      />
                    </div>
                  )
                }
              >
                <div
                  style={{
                    fontSize: 10,
                    color: "gray",
                    fontWeight: "normal",
                    whiteSpace: "nowrap",
                    overflow: "hidden",
                    textOverflow: "ellipsis",
                    cursor: "pointer",
                  }}
                >
                  {item.description ? (
                    item.description
                  ) : (
                    <div
                      className={css`
                        opacity: ${popoverVisible ? 1 : 0};
                        transition: all 0.2s ease;
                        &:hover {
                          opacity: 1;
                        }
                      `}
                    >
                      暂无描述，点击添加
                    </div>
                  )}
                </div>
              </Popover>
            </Col>
            <Col span={8} style={{ textAlign: "right" }}>
              <Space style={{ marginRight: 4 }}>
                <NamespacePrivate reload={reload} item={item} />
                <CpuMemory title="空间资源总使用量" namespaceID={item.id} />
                <ServiceEndpoint namespaceId={item.id} />
              </Space>
            </Col>
          </Row>
        </div>
      }
      bordered={false}
    >
      <Spin spinning={deleting || loading}>
        <Row gutter={[8, 8]}>
          {item.projects?.map((data) => (
            <Col key={data.id} md={12} xs={24} sm={24}>
              <ProjectDetail
                namespace={item.name}
                namespaceId={item.id}
                item={data}
              />
            </Col>
          ))}

          <Col md={12} xs={24} sm={24}>
            <CreateProjectModal namespaceId={item.id} />
          </Col>
        </Row>
      </Spin>
    </Card>
  );
};

export default memo(Item);

const TitleNamespace = styled.div`
  font-size: 12px;
  font-weight: normal;
`;

const TitleNamespaceName = styled.span`
  font-family: "Gill Sans", "Gill Sans MT", Calibri, "Trebuchet MS", sans-serif;
  font-weight: 500;
  font-size: 18px;
  margin-left: 3px;
`;

const useIsOwned = (item: components["schemas"]["types.NamespaceModel"]) => {
  const { user, isAdmin } = useAuth();
  const isOwned = useCallback(() => {
    return isAdmin() || item.creatorEmail === user.email;
  }, [isAdmin, item.creatorEmail, user.email]);

  return isOwned();
};

const NamespacePrivate: React.FC<{
  item: components["schemas"]["types.NamespaceModel"];
  reload: () => void;
}> = ({ item, reload }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isTransferOpen, setIsTransferOpen] = useState(false);
  const [transferEmail, setTransferEmail] = useState("");

  const initValues = yaml.dump(item.members.map((v) => v.email));
  const [value, setValue] = useState(initValues);

  if (!useIsOwned(item)) {
    return (
      <Tooltip
        placement="top"
        overlayStyle={{ fontSize: 12 }}
        title={`此项目管理员是: ${item.creatorEmail}`}
      >
        {!item.private ? <UnlockOutlined /> : <LockOutlined />}
      </Tooltip>
    );
  }

  return (
    <Space>
      <Modal
        width={"60%"}
        destroyOnClose
        title="修改空间成员"
        open={isModalOpen}
        onOk={() => {
          try {
            let v = yaml.load(value);
            ajax
              .POST("/api/namespaces/sync_members", {
                body: {
                  id: item.id,
                  emails: (v as string[]) || [],
                },
              })
              .then(({ error }) => {
                if (error) {
                  message.error(error.message);
                  return;
                }
                message.success("修改成功");
                reload();
                setIsModalOpen(false);
              });
          } catch (e) {
            message.error("yaml 格式不正确");
          }
        }}
        onCancel={() => {
          setIsModalOpen(false);
        }}
      >
        <Form>
          <Row style={{ marginBottom: 5 }}>
            <Col span={24}>
              <Alert
                message={
                  <div>
                    <div>参考下面格式👇，自行修改，输入用户邮箱</div>
                    <div>- a@qq.com</div>
                    <div>- v@qq.com</div>
                  </div>
                }
                type="success"
              />
            </Col>
          </Row>
          <Row>
            <Col span={12}>
              <MyCodeMirror
                value={value}
                onChange={(v) => setValue(v)}
                mode="yaml"
              />
            </Col>
            <Col span={12}>
              <DiffViewer
                mode={"yaml"}
                styles={{
                  line: { fontSize: 12, lineHeight: 10 },
                  gutter: { padding: "0 5px", minWidth: 20 },
                  marker: { padding: "0 6px" },
                  diffContainer: {
                    height: "100%",
                    display: "block",
                    width: "100%",
                    overflowX: "auto",
                  },
                }}
                showDiffOnly={false}
                oldValue={initValues}
                newValue={value}
                splitView={false}
              />
            </Col>
          </Row>
        </Form>
      </Modal>
      <Tooltip
        overlayStyle={{ fontSize: 12 }}
        trigger={["hover"]}
        title={<div>空间管理员邮箱: {item.creatorEmail}</div>}
      >
        <IconFont style={{ cursor: "pointer" }} name="#icon-crown" />
      </Tooltip>
      <Popconfirm
        overlayClassName="fullwidthpop"
        title="转让空间"
        open={isTransferOpen}
        description={
          <Input
            type="email"
            value={transferEmail}
            onChange={(v) => {
              setTransferEmail(v.target.value);
            }}
            style={{ width: "80%", fontSize: 12 }}
            placeholder="新空间管理员邮箱"
          />
        }
        onConfirm={() => {
          ajax
            .POST("/api/namespaces/transfer", {
              body: { newAdminEmail: transferEmail, id: item.id },
            })
            .then(({ error }) => {
              if (error) {
                message.error(error.message);
                return;
              }
              message.success("操作成功");
              setIsTransferOpen(false);
              reload();
              setTransferEmail("");
            });
        }}
        onCancel={() => {
          setIsTransferOpen(false);
          setTransferEmail("");
        }}
        overlayInnerStyle={{ width: 500 }}
        okText="确定"
        cancelText="取消"
      >
        <Tooltip overlayStyle={{ fontSize: 12 }} title="转让空间">
          <IconFont
            style={{ cursor: "pointer" }}
            onClick={() => setIsTransferOpen(true)}
            name="#icon-zhuanyi"
          />
        </Tooltip>
      </Popconfirm>
      <Popconfirm
        title="修改访问权限"
        description={`确定要修改成 ${item.private ? "public" : "private"} 吗`}
        onConfirm={() => {
          ajax
            .POST("/api/namespaces/update_private", {
              body: { id: item.id, private: !item.private },
            })
            .then(() => {
              message.success("修改成功");
              reload();
            });
        }}
        okText="Yes"
        cancelText="No"
      >
        <Tooltip overlayStyle={{ fontSize: 12 }} title="修改空间访问权限">
          {!item.private ? <UnlockOutlined /> : <LockOutlined />}
        </Tooltip>
      </Popconfirm>
      {item.private && (
        <Tooltip overlayStyle={{ fontSize: 12 }} title="成员管理">
          <IconFont
            name="#icon-chengyuanguanli_huaban1"
            style={{ cursor: "pointer" }}
            onClick={() => setIsModalOpen(true)}
          />
        </Tooltip>
      )}
    </Space>
  );
};
