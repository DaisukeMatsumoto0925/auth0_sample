import logo from "./logo.svg";
import "./App.css";
import { useAuth0 } from "@auth0/auth0-react";

function App() {
  const { loginWithRedirect, isAuthenticated } = useAuth0();

  const onClickLogin = () => {
    loginWithRedirect();
  };

  return (
    <div className="App">
      <div
        style={{ display: "flex", justifyContent: "center", marginTop: "64px" }}
      >
        <button onClick={onClickLogin} disabled={isAuthenticated}>
          {isAuthenticated ? "ログイン済み" : "ログイン"}
        </button>
      </div>
    </div>
  );
}

export default App;
