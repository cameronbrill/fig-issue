import { createContext, useContext, useEffect, useState } from "react";
import { User } from "../../types/user";
import { Session } from "@supabase/supabase-js";
import { supabase } from "../../lib/supabase/client";

type AuthContextState = {
  user: User | undefined;
  session: Session | undefined;
};

const AuthContext = createContext<AuthContextState>({
  user: undefined,
  session: undefined,
});

type AuthProviderProps = {
  children: React.ReactNode;
};

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [session, setSession] = useState<Session | undefined>(
    supabase.auth.session() ?? undefined
  );
  const [user, setUser] = useState<User | undefined>(
    session?.user ?? undefined
  );

  useEffect(() => {
    const { data: authListener } = supabase.auth.onAuthStateChange(
      async (_, session) => {
        setSession(session ?? undefined);
        setUser(session?.user ?? undefined);
      }
    );

    return () => {
      authListener!.unsubscribe();
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const value = {
    user: user || undefined,
    session: session || undefined,
  };
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuthContext = (): AuthContextState => {
  const context = useContext(AuthContext);
  return context;
};
