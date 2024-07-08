import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { Layout } from "../pages/Layout";
import { Puzzle } from "../components/Puzzle";
import { BaseErrorBoundary } from "../components/BaseErrorBoundary";
import { Home } from "../pages/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    errorElement: <BaseErrorBoundary />,
    children: [
      { index: true, element: <Home /> },
      { path: "/puzzle", element: <Puzzle /> },
    ],
  },
]);

export const PublicRoutes = () => {
  return <RouterProvider router={router} />;
};
