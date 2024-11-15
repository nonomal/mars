import {
  CLEAR_CREATE_PROJECT_LOG,
  SET_DEPLOY_STATUS,
  SET_PROCESS_PERCENT,
  APPEND_CREATE_PROJECT_LOG,
  SET_CREATE_PROJECT_LOADING,
  CLEAN_PROJECT,
} from "../actionTypes";

import { set, get } from "lodash";
import pb from "../../api/websocket";

export enum DeployStatus {
  DeployUnknown = "unknown",
  DeployFailed = "failed",
  DeployCanceled = "canceled",
  DeploySuccess = "success",
}

export interface Output {
  type: pb.websocket.ResultType;
  log: string;
  containers?: pb.websocket.Container[];
}

export interface CreateProjectItem {
  isLoading: boolean;
  deployStatus: DeployStatus;
  output: Output[];
  processPercent: number;
}

export const selectList = (state: { createProject: List }): List =>
  state.createProject;

export interface List {
  [id: string]: CreateProjectItem;
}

const initialState: List = {};

export default function createProject(
  state = initialState,
  action: {
    type: string;
    data?: {
      id: string;
      isLoading: boolean;
      output: Output;
      deployStatus: string;
      processPercent: number;
    };
  },
) {
  switch (action.type) {
    case CLEAN_PROJECT:
      if (action.data) {
        delete state[action.data.id];
        return {
          ...state,
        };
      }

      return state;
    case SET_PROCESS_PERCENT:
      if (action.data) {
        return {
          ...set(
            state,
            [action.data.id, "processPercent"],
            action.data.processPercent,
          ),
        };
      }

      return state;
    case SET_DEPLOY_STATUS:
      if (action.data) {
        return {
          ...set(
            state,
            [action.data.id, "deployStatus"],
            action.data.deployStatus,
          ),
        };
      }

      return state;
    case SET_CREATE_PROJECT_LOADING:
      if (action.data) {
        console.log("action.data.isLoading", action.data.isLoading);
        return {
          ...set(state, [action.data.id, "isLoading"], action.data.isLoading),
        };
      }

      return state;
    case CLEAR_CREATE_PROJECT_LOG:
      if (action.data) {
        return { ...set(state, [action.data.id, "output"], []) };
      }

      return state;
    case APPEND_CREATE_PROJECT_LOG:
      if (action.data) {
        let g = get(state, [action.data.id, "output"], []);
        return {
          ...set(state, [action.data.id, "output"], [...g, action.data.output]),
        };
      }

      return state;
    default:
      return state;
  }
}
