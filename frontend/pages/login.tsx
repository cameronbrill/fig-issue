import { NextPage } from "next/types";
import { Login as AuthLogin } from "../components/Auth/Login";

const Login: NextPage = () => {
  return (
    <section>
      <AuthLogin />
    </section>
  );
};

export default Login;
