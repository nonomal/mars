import React, { memo, useState } from "react";
import { getServiceEndpoints } from "../api/namespace";
import { Popover } from "antd";

const ServiceEndpoint: React.FC<{ namespaceId: number; projectName?: string }> =
  ({ namespaceId, projectName }) => {
    const [endpoints, setEndpoints] = useState<{ [name: string]: string[] }>(
      {}
    );

    return (
      <Popover
        placement="right"
        title={"链接"}
        content={Object.entries(endpoints).map(([k, v]) =>
          v.map((link) => (
            <div key={link} onClick={(e) => e.stopPropagation()}>
              <span style={{ marginRight: 5 }}>{k}:</span>
              <a href={link} target="_blank">
                {link}
              </a>
            </div>
          ))
        )}
        trigger="hover"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          style={{ width: 18, height: 18, flexShrink: 0 }}
          stroke="currentColor"
          onMouseEnter={(e) => {
            getServiceEndpoints(namespaceId, projectName).then((res) => {
              setEndpoints(res.data.data);
            });
          }}
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
          />
        </svg>
      </Popover>
    );
  };

export default memo(ServiceEndpoint);