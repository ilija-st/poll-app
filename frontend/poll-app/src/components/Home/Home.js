import React from "react";
import { useOutletContext } from "react-router-dom";

const Home = () => {
  const { user } = useOutletContext();
  const { jwtToken } = useOutletContext();

  return (
    <>
      <div className="text-center">
        {jwtToken !== "" ? (
          <h2>
            Welcome {user.first_name}, are you ready to create or vote on your
            first poll? :)
          </h2>
        ) : (
          <h2>Login or register to create or vote on your first poll! :)</h2>
        )}
      </div>
    </>
  );
};

export default Home;
