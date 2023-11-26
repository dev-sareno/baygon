import Facade from "./pages/Facade/Facade";
import Header from "./components/Header/Header";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import PrivacyPolicy from "./pages/PrivacyPolicy/PrivacyPolicy";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Facade />,
  },
  {
    path: "/privacy",
    element: <PrivacyPolicy />,
  },
]);

function App() {
  return (
    <>
      <Header />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
