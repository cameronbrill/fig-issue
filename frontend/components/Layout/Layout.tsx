import { Auth } from "@supabase/ui";
import { Props as AuthProps } from "@supabase/ui/dist/cjs/components/Auth/Auth";
import Head from "next/head";
import React, { ReactNode, useEffect, useState } from "react";
import useSWR from "swr";
import { supabase } from "../../lib/supabase/client";
import { useAuthContext } from "../Auth/UserContext";

type LayoutProps = {
  children: ReactNode;
};

const fetcher = (url: string, token: string) =>
  fetch(url, {
    method: "GET",
    headers: new Headers({ "Content-Type": "application/json", token }),
    credentials: "same-origin",
  }).then((res) => res.json());

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const [authView, setAuthView] = useState<AuthProps["view"]>("sign_in");
  const { user } = useAuthContext();
  useEffect(() => {
    const { data: authListener } = supabase.auth.onAuthStateChange(
      (event, session) => {
        if (event === "PASSWORD_RECOVERY") setAuthView("update_password");
        if (event === "USER_UPDATED")
          setTimeout(() => setAuthView("sign_in"), 1000);
        // Send session to /api/auth route to set the auth cookie.
        // NOTE: this is only needed if you're doing SSR (getServerSideProps)!
        fetch("/api/auth", {
          method: "POST",
          headers: new Headers({ "Content-Type": "application/json" }),
          credentials: "same-origin",
          body: JSON.stringify({ event, session }),
        }).then((res) => res.json());
      }
    );

    return () => {
      authListener?.unsubscribe();
    };
  }, []);
  return (
    <div>
      <Head>
        <title>Fig Issue</title>
      </Head>
      <header></header>
      <main>
        {(user && children) || (
          <Auth
            supabaseClient={supabase}
            providers={["google"]}
            view={authView}
            socialLayout="horizontal"
            socialButtonSize="xlarge"
          />
        )}
      </main>
      <footer></footer>
    </div>
  );
};

export default Layout;
