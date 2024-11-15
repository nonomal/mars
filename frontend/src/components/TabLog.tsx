import React, { memo, useEffect, useState, useCallback, useMemo } from "react";
import { Radio, Skeleton, Button, message, Empty } from "antd";
import LazyLog from "../pkg/lazylog/components/LazyLog";
import { getToken } from "./../utils/token";
import { useSelector } from "react-redux";
import { selectPodEventProjectID } from "../store/reducers/podEventWatcher";
import { debounce } from "lodash";
import PodStateTag from "./PodStateTag";
import { css } from "@emotion/css";
import ajax from "../api/ajax";
import { components } from "../api/schema";

const ProjectContainerLogs: React.FC<{
  id: number;
  namespace: string;
}> = ({ id, namespace }) => {
  console.log("render: TabLog");
  const [value, setValue] = useState<string>("");
  const [list, setList] = useState<
    components["schemas"]["types.StateContainer"][]
  >([]);

  const listContainer = useCallback(async () => {
    return ajax
      .GET("/api/projects/{id}/containers", { params: { path: { id: id } } })
      .then(({ data, error }) => {
        if (error) {
          return;
        }
        setList(data.items);
        return data;
      });
  }, [id]);
  useEffect(() => {
    if (
      list.length > 0 &&
      !list.map((v) => v.pod + "|" + v.container).includes(value)
    ) {
      setValue(list[0].pod + "|" + list[0].container);
    }
  }, [list, value]);
  const projectIDStr = useSelector(selectPodEventProjectID);
  useEffect(() => {
    let d = debounce(() => {
      listContainer();
    }, 1000);
    console.log("ns event: ", projectIDStr);
    if (projectIDStr.split("-").length === 2) {
      let pid = Number(projectIDStr.split("-")[1]);
      if (pid === Number(id)) {
        d();
      }
    }
    return () => {
      d.cancel();
    };
  }, [projectIDStr, listContainer, id]);

  useEffect(() => {
    listContainer().then((data) => {
      if (!data) {
        return;
      }
      if (data.items.length > 0) {
        let first = data.items[0];
        setValue(first.pod + "|" + first.container);
      }
    });
  }, [setList, id, namespace, listContainer]);

  const [timestamp, setTimestamp] = useState(new Date().getTime());
  let [pod, container] = useMemo(() => (value as string).split("|"), [value]);

  const reloadLog = useCallback((e: any) => {
    setValue(e.target.value);
    setTimestamp(new Date().getTime());
  }, []);

  return (
    <div style={{ display: "flex", flexDirection: "column", height: "100%" }}>
      {list.length > 0 ? (
        <>
          <Radio.Group value={value} style={{ marginBottom: 10 }}>
            {list.map((item) => (
              <Radio
                onClick={reloadLog}
                key={item.pod + "|" + item.container}
                value={item.pod + "|" + item.container}
              >
                {item.container}
                <PodStateTag pod={item} />
              </Radio>
            ))}
          </Radio.Group>

          <div
            className={css`
              pre {
                margin: 0;
              }
            `}
            style={{
              fontFamily: '"Fira code", "Fira Mono", monospace',
              fontSize: 12,
              height: "100%",
            }}
          >
            <Skeleton active loading={!(pod && container)}>
              <MyLogUtil
                namespace={namespace}
                freshTime={timestamp}
                pod={pod}
                container={container}
                onError={{
                  praseJsonError: () => {
                    return (
                      <span style={{ textAlign: "center" }}>
                        <Button
                          type="text"
                          style={{ color: "red", fontSize: 12 }}
                          size="small"
                          onClick={() => setTimestamp(new Date().getTime())}
                        >
                          点击重新加载
                        </Button>
                      </span>
                    );
                  },
                  on404Error: (e) => {
                    listContainer().then((data) => {
                      if (!data) {
                        return;
                      }
                      if (data.items.length > 0) {
                        let first = data.items[0];
                        setValue(first.pod + "|" + first.container);
                      }
                    });
                  },
                }}
              />
            </Skeleton>
          </div>
        </>
      ) : (
        <div
          style={{
            height: "100%",
            width: "100%",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Empty description="列表还没有任何容器" />
        </div>
      )}
    </div>
  );
};

const MyLogUtil: React.FC<{
  namespace: string;
  pod: string;
  container: string;
  freshTime: number;
  onError?: {
    praseJsonError: (e: any) => any;
    on404Error: (e: any) => void;
  };
}> = ({ namespace, pod, container, onError, freshTime }) => {
  const getUrl = useCallback(() => {
    return `${process.env.REACT_APP_BASE_URL}/api/containers/namespaces/${namespace}/pods/${pod}/containers/${container}/stream_logs?timestamp=${freshTime}&show_events=1`;
  }, [namespace, pod, container, freshTime]);

  return (
    <LazyLog
      renderErrLineFunc={(e: any) => {
        return JSON.parse(e.body).error.message;
      }}
      fetchOptions={{ headers: { Authorization: getToken() } }}
      enableSearch
      selectableLines
      captureHotkeys
      formatPart={(text: string) => {
        // console.log(text);
        let res = JSON.parse(text);
        if (res.error) {
          if (onError?.praseJsonError) {
            return onError?.praseJsonError(res.error);
          }
          return res.error;
        }
        return res.result.log;
      }}
      stream
      onError={(e: any) => {
        if (e.status === 404) {
          message.error(JSON.parse(e.body).error.message);
          onError?.on404Error(e);
        }
      }}
      follow={true}
      url={getUrl()}
    />
  );
};

export const LogUtil = memo(MyLogUtil);
export default memo(ProjectContainerLogs);
