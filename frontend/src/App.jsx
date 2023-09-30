import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { useAppState } from "./context/AppStateContext";
import { GetState } from "../wailsjs/go/main/App";
import { EventsOn } from "../wailsjs/runtime/runtime";

export default function App() {
  const { state, dispatch } = useAppState();
  const navigate = useNavigate();

  useEffect(() => {
    async function loadData() {
      const state = await GetState();
      dispatch({ type: "INIT", payload: state });
    }
    loadData();

    EventsOn("stateChanged", (state) => {
      console.log("stateChanged", state);
      dispatch({ type: "UPDATE", payload: state });
    });

    EventsOn("routeChanged", (route) => {
      console.log("routeChanged", route);
      navigate(route);
    });

    return () => {};
  }, []);

  return (
    <div className="App">
      {state.isLoading && <h2>Loading...</h2>}
      {state.isLoading === false && <Outlet />}
    </div>
  );
}
