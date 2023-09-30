import React from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./theme/main.scss";
import App from "./App";
import { AppStateProvider } from "./context/AppStateContext";
import ErrorPage from "./error-page";
import HomeScreen from "./screen/Home.jsx";
import AboutScreen from "./screen/About";
import GuideScreen from "./screen/Guide";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/",
        element: <HomeScreen />,
      },
      {
        path: "/about",
        element: <AboutScreen />,
      },
      {
        path: "/guide",
        element: <GuideScreen />,
      },
    ],
  },
]);

const root = createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <AppStateProvider>
      <RouterProvider router={router} />
    </AppStateProvider>
  </React.StrictMode>
);
