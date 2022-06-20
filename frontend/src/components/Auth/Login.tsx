import { supabaseClient } from "@supabase/supabase-auth-helpers/nextjs";
import { Auth } from "@supabase/ui";
import { useUser } from "@supabase/supabase-auth-helpers/react";

export const Login = () => {
  const { user, error } = useUser();

  if (!user)
    return (
      <div>
        {error && <p>{error.message}</p>}
        <div>
          <Auth
            supabaseClient={supabaseClient}
            providers={["google"]}
            socialLayout="horizontal"
            socialButtonSize="xlarge"
          />
        </div>
      </div>
    );

  return (
    <>
      <button onClick={() => supabaseClient.auth.signOut()}>Sign out</button>
      <p>user:</p>
      <pre>{JSON.stringify(user, null, 2)}</pre>
    </>
  );
};
