import { NextPage } from "next/types";
import Profile from "../components/Profile/Profile";
import { withPageAuth } from "@supabase/supabase-auth-helpers/nextjs";
import { User } from "../types/user";

type MeProps = {
  user: User;
};

const Me: NextPage<MeProps> = ({ user }) => {
  return (
    <section>
      <div style={{ maxWidth: "420px", margin: "96px auto" }}>
        <Profile user={user} />
      </div>
    </section>
  );
};

export default Me;

export const getServerSideProps = withPageAuth({ redirectTo: "/login" });
