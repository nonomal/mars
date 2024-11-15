import React, { useCallback, useState, Fragment, memo, useMemo } from "react";
import { Form, Input, InputNumber, Radio, Select, Switch } from "antd";
import { omitEqual } from "../../utils/obj";
import { css } from "@emotion/css";
import { components, MarsElementType } from "../../api/schema.d";
import { set, findIndex } from "lodash";

const Option = Select.Option;
const { TextArea } = Input;

interface st {
  textarea?: React.CSSProperties;
  input?: React.CSSProperties;
  inputNumber?: React.CSSProperties;
  label?: React.CSSProperties;
  formItem?: React.CSSProperties;
  radioGroup?: React.CSSProperties;
  radio?: React.CSSProperties;
  select?: React.CSSProperties;
  selectOption?: React.CSSProperties;
  switch?: React.CSSProperties;
}
const isTrue = (v: any): boolean => {
  switch (v) {
    case 1:
    case true:
    case "1":
    case "True":
    case "true":
      return true;
    default:
      return false;
  }
};
const initStyle = {
  input: {},
  label: {},
  formItem: {},
  radioGroup: {},
  radio: {},
  select: {},
  selectOption: {},
  switch: {},
};

const Elements: React.FC<{
  value?: components["schemas"]["websocket.ExtraValue"][];
  onChange?: (value: components["schemas"]["websocket.ExtraValue"][]) => void;
  elements: components["schemas"]["mars.Element"][];
  style?: st;
  id?: string;
}> = ({ elements, style, value, onChange, id }) => {
  const getElement = useCallback(
    (
      item: components["schemas"]["websocket.ExtraValue"],
      ele: components["schemas"]["mars.Element"][],
    ): React.ReactNode => {
      for (let i = 0; i < ele.length; i++) {
        let ev = ele[i];
        if (ev.path === item.path) {
          return (
            <Element
              value={item.value}
              onChange={(changeValue) => {
                let tmp: components["schemas"]["websocket.ExtraValue"][] =
                  value || [];
                let idx = findIndex(tmp, (v) => v.path === item.path);
                set(tmp, idx, {
                  path: item.path,
                  value: String(changeValue),
                });
                onChange?.(tmp);
                return tmp;
              }}
              element={ev}
              style={style ? style : initStyle}
            />
          );
        }
      }
      return <></>;
    },
    [onChange, style, value],
  );

  let initValues = useMemo(() => {
    return elements
      ? elements.map((item): components["schemas"]["websocket.ExtraValue"] => {
          let itemValue: any = item.default;
          if (!!value) {
            for (let i = 0; i < value.length; i++) {
              if (value[i].path === item.path) {
                itemValue = value[i].value;
                if (item.type === MarsElementType.ElementTypeSwitch) {
                  itemValue = isTrue(itemValue);
                }
                if (item.type === MarsElementType.ElementTypeInputNumber) {
                  itemValue = Number(itemValue);
                }
                break;
              }
            }
          }
          return { path: item.path, value: itemValue };
        })
      : [];
  }, [elements, value]);

  return (
    <div style={{ width: "100%" }} id={id}>
      {initValues &&
        initValues.map((item) => (
          <Fragment key={item.path}>{getElement(item, elements)}</Fragment>
        ))}
    </div>
  );
};

const Element: React.FC<{
  value: any;
  onChange: (v: any) => void;
  element: components["schemas"]["mars.Element"];
  style: st;
}> = ({ element, style, value: v, onChange }) => {
  const [value, setValue] = useState(v);
  switch (element.type) {
    case MarsElementType.ElementTypeInput:
      return (
        <Form.Item
          label={<div style={style.label}>{element.description}</div>}
          style={style.formItem}
        >
          <Input
            defaultValue={element.default}
            value={value}
            onChange={(e) => {
              setValue(e.target.value);
              onChange(e.target.value);
            }}
            style={style.input}
          />
        </Form.Item>
      );
    case MarsElementType.ElementTypeTextArea:
      return (
        <Form.Item
          className={css`
            margin-bottom: 10px;
            .ant-form-item-row {
              display: block;
            }
          `}
          labelAlign={"left"}
          label={<div style={style.label}>{element.description}</div>}
          style={{ width: "100%" }}
        >
          <TextArea
            defaultValue={element.default}
            value={value}
            onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => {
              setValue(e.target.value);
              onChange(e.target.value);
            }}
            style={style.textarea}
          />
        </Form.Item>
      );
    case MarsElementType.ElementTypeInputNumber:
      return (
        <Form.Item
          label={<div style={style.label}>{element.description}</div>}
          style={style.formItem}
        >
          <InputNumber
            defaultValue={Number(element.default)}
            style={style.inputNumber}
            value={value}
            onChange={(e) => {
              setValue(e);
              onChange(e);
            }}
          />
        </Form.Item>
      );
    case MarsElementType.ElementTypeRadio:
    case MarsElementType.ElementTypeNumberRadio:
      return (
        <Form.Item
          label={<div style={style.label}>{element.description}</div>}
          style={style.formItem}
        >
          <Radio.Group
            defaultValue={element.default}
            style={style.radioGroup}
            value={value}
            onChange={(e) => {
              setValue(e.target.value);
              onChange(e.target.value);
            }}
          >
            {element.selectValues.map((i, k) => (
              <Radio key={k} value={i} style={style.radio}>
                {i}
              </Radio>
            ))}
          </Radio.Group>
        </Form.Item>
      );
    case MarsElementType.ElementTypeSelect:
    case MarsElementType.ElementTypeNumberSelect:
      return (
        <Form.Item
          label={<div style={style.label}>{element.description}</div>}
          style={style.formItem}
        >
          <Select
            defaultValue={element.default}
            style={style.select}
            value={value}
            onChange={(e) => {
              setValue(e);
              onChange(e);
            }}
          >
            {element.selectValues.map((i, k) => (
              <Option value={i} key={k}>
                <div style={style.selectOption}>{i}</div>
              </Option>
            ))}
          </Select>
        </Form.Item>
      );
    case MarsElementType.ElementTypeSwitch:
      return (
        <Form.Item
          label={<div style={style.label}>{element.description}</div>}
          style={style.formItem}
        >
          <Switch
            defaultChecked={isTrue(element.default)}
            style={style.switch}
            checked={isTrue(value)}
            onChange={(e) => {
              setValue(e);
              onChange(e);
            }}
          />
        </Form.Item>
      );
    default:
      return <></>;
  }
};

export default memo(Elements, (prev, next) =>
  omitEqual(prev, next, "onChange"),
);
