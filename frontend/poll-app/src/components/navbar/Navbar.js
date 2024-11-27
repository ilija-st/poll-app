import * as React from "react";
import { styled, alpha } from "@mui/material/styles";
import Box from "@mui/material/Box";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import { useNavigate } from "react-router-dom";

const StyledToolbar = styled(Toolbar)(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  justifyContent: "space-between",
  flexShrink: 0,
  borderRadius: `calc(${theme.shape.borderRadius}px + 8px)`,
  backdropFilter: "blur(24px)",
  border: "1px solid",
  borderColor: (theme.vars || theme).palette.divider,
  backgroundColor: theme.vars
    ? `rgba(${theme.vars.palette.background.defaultChannel} / 0.4)`
    : alpha(theme.palette.background.default, 0.4),
  boxShadow: (theme.vars || theme).shadows[1],
  padding: "8px 12px",
}));

export default function Navbar({
  jwtToken,
  updateJwtToken,
  removeUser,
  toggleRefresh,
}) {
  const navigate = useNavigate();

  const goToSignIn = () => {
    navigate("/signin");
  };

  const goToSignUp = () => {
    navigate("/signup");
  };

  const navigateHome = () => {
    navigate("/");
  };

  const navigatePolls = () => {
    navigate("/polls");
  };

  const navigateCreatePoll = () => {
    navigate("/polls/new");
  };

  const logout = () => {
    const requestOptions = {
      method: "GET",
      credentials: "include",
    };

    fetch(`${process.env.REACT_APP_BACKEND}/logout`, requestOptions)
      .catch((error) => {
        console.log("error logging out", error);
      })
      .finally(() => {
        updateJwtToken("");
        removeUser();
        toggleRefresh(false);
        navigate("/signin");
      });

    navigate("/signin");
  };

  return (
    <AppBar
      position="fixed"
      enableColorOnDark
      sx={{
        boxShadow: 0,
        bgcolor: "transparent",
        backgroundImage: "none",
        mt: "calc(var(--template-frame-height, 0px) + 28px)",
      }}
    >
      <Container maxWidth="lg">
        <StyledToolbar variant="dense" disableGutters>
          <Box
            sx={{ flexGrow: 1, display: "flex", alignItems: "center", px: 0 }}
          >
            <Box sx={{ display: { xs: "none", md: "flex" } }}>
              <Button
                variant="text"
                color="info"
                size="small"
                onClick={navigateHome}
              >
                Home
              </Button>
              {jwtToken !== "" && (
                <Button
                  variant="text"
                  color="info"
                  size="small"
                  onClick={navigatePolls}
                >
                  Polls
                </Button>
              )}
              {jwtToken !== "" && (
                <Button
                  variant="text"
                  color="info"
                  size="small"
                  onClick={navigateCreatePoll}
                >
                  Create a Poll
                </Button>
              )}
            </Box>
          </Box>
          {jwtToken === "" ? (
            <Box
              sx={{
                display: { xs: "none", md: "flex" },
                gap: 1,
                alignItems: "center",
              }}
            >
              <Button
                onClick={goToSignIn}
                color="primary"
                variant="text"
                size="small"
              >
                Sign in
              </Button>
              <Button
                onClick={goToSignUp}
                color="primary"
                variant="contained"
                size="small"
              >
                Sign up
              </Button>
            </Box>
          ) : (
            <Box
              sx={{
                display: { xs: "none", md: "flex" },
                gap: 1,
                alignItems: "center",
              }}
            >
              <Button
                onClick={logout}
                color="primary"
                variant="text"
                size="small"
              >
                Logout
              </Button>
            </Box>
          )}
        </StyledToolbar>
      </Container>
    </AppBar>
  );
}
