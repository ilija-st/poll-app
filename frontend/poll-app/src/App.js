import { useCallback, useEffect, useState } from "react";
import Navbar from "./components/navbar/Navbar";
import AppTheme from "./components/shared-theme/AppTheme";
import { CssBaseline } from "@mui/material";
import Container from "@mui/material/Container";
import { Outlet } from "react-router-dom";

function App() {
  const [jwtToken, setJwtToken] = useState("");
  const [user, setUser] = useState({});
  const [tickInterval, setTickInterval] = useState();

  const toggleRefresh = useCallback(
    (status) => {
      console.log("clicked");

      if (status) {
        console.log("turning on ticking");
        let i = setInterval(() => {
          const requestOptions = {
            method: "GET",
            credentials: "include",
          };

          fetch(`/refresh`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
              if (data.access_token) {
                setJwtToken(data.access_token);
              }
            })
            .catch((error) => {
              console.log("user is not logged in");
              setUser({});
            });
        }, 600000);
        setTickInterval(i);
        console.log("setting tick interval to", i);
      } else {
        console.log("turning off ticking");
        console.log("turning off tickInterval", tickInterval);
        setTickInterval(null);
        clearInterval(tickInterval);
      }
    },
    [tickInterval]
  );

  useEffect(() => {
    if (jwtToken === "") {
      const requestOptions = {
        method: "GET",
        credentials: "include",
      };

      fetch(`/refresh`, requestOptions)
        .then((response) => response.json())
        .then((data) => {
          if (data.access_token) {
            setJwtToken(data.access_token);
            toggleRefresh(true);
          }
        })
        .catch((error) => {
          console.log("user is not logged in", error);
        });
    }
  }, [jwtToken, toggleRefresh]);

  return (
    <AppTheme>
      <CssBaseline enableColorScheme />
      <Container
        maxWidth="lg"
        component="main"
        sx={{ display: "flex", flexDirection: "column", my: 16, gap: 4 }}
      >
        <Navbar
          jwtToken={jwtToken}
          updateJwtToken={(val) => setJwtToken(val)}
          removeUser={() => setUser({})}
          toggleRefresh={() => toggleRefresh()}
        />
        <Outlet
          context={{
            jwtToken,
            user,
            setUser,
            setJwtToken,
            toggleRefresh,
          }}
        />
      </Container>
    </AppTheme>
  );
}

export default App;
