export const initialAppState = {
    isLoading: true,
    nodes: [],
    nodeIndex: 0,
    bufferIndex: 0,
}

export const reducer = (state, action) => {
  switch (action.type) {
    case "INIT":
      return {...state, ...action.payload, isLoading: false};
    case "UPDATE":
      return {...state, ...action.payload};
    default:
      return state;
  }
};
