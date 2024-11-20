import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import SignIn from "./components/sign-in/SignIn";
import SignUp from "./components/sign-up/SignUp";
import ErrorPage from "./components/error-page/ErrorPage";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Polls from "./components/poll/Polls";
import Home from "./components/Home/Home";
import CreatePoll from "./components/poll/CreatePoll";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Home /> },
      {
        path: "/signin",
        element: <SignIn />,
      },
      {
        path: "/signup",
        element: <SignUp />,
      },
      {
        path: "/polls",
        element: <Polls />,
      },
      {
        path: "/polls/new",
        element: <CreatePoll />,
      },
    ],
  },
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
