import "../styles/globals.css";
import "@highlight-run/react/dist/highlight.css";
import type { AppProps } from "next/app";
import { H } from "highlight.run";
import Layout from "../components/Layout/Layout";
import { ErrorBoundary } from "@highlight-run/react";
import { HighlightOptions } from "../lib/highlight.config";
import { UserProvider } from "@supabase/supabase-auth-helpers/react";
import { supabaseClient } from "@supabase/supabase-auth-helpers/nextjs";

H.init(process.env.NEXT_PUBLIC_HIGHLIGHT_PROJECT_ID, HighlightOptions);

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ErrorBoundary>
      <Layout>
        <UserProvider supabaseClient={supabaseClient}>
          <Component {...pageProps} />
        </UserProvider>
      </Layout>
    </ErrorBoundary>
  );
}

export default MyApp;
