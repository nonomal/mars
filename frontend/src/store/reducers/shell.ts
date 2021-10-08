import { SET_SHELL_SESSION_ID, SET_SHELL_LOG } from "./../actionTypes";
const initialState: {
  [id: string]: { sessionID: string; log: string; logCount: number };
} = {};

export const selectSessions = (state: {
  shell: { [id: string]: { sessionID: string; log: string; logCount: number } };
}) => state.shell;

export default function shell(
  state = initialState,
  action: { type: string; data: { id: string; sessionID: string; log: string } }
) {
  switch (action.type) {
    case SET_SHELL_LOG:
        let count = 0
        if (state[action.data.id] && state[action.data.id].logCount) {
            count = state[action.data.id].logCount
        }
      return {
        ...state,
        [action.data.id]: {
          ...state[action.data.id],
          log: action.data.log,
          logCount: count + 1,
        },
      };
    case SET_SHELL_SESSION_ID:
      console.log("SET_SHELL_SESSION_ID", {
        ...state,
        [action.data.id]: { sessionID: action.data?.sessionID },
      });
      return {
        ...state,
        [action.data.id]: { sessionID: action.data?.sessionID },
      };
    default:
      return state;
  }
}