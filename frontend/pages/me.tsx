import { NextPage } from "next/types";
import Profile from "../components/Profile/Profile";
import { useAuthContext } from "../components/Auth/UserContext";

const Me: NextPage = () => {
  const { user } = useAuthContext();
  return (
    <section>
      <div style={{ maxWidth: "420px", margin: "96px auto" }}>
        <Profile user={user} />
      </div>
    </section>
  );
};

export default Me;
