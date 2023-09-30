import React, { createContext, useContext, useReducer } from "react";
import { reducer, initialAppState } from "./reducers.js";

const AppStateContext = createContext(null);

export const AppStateProvider = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, initialAppState);

  return (
    <AppStateContext.Provider value={{ state, dispatch }}>
      {children}
    </AppStateContext.Provider>
  );
};

export const useAppState = () => {
  const context = useContext(AppStateContext);
  if (context === undefined) {
    throw new Error("useAppState must be used within a AppStateProvider");
  }
  return context;
};
