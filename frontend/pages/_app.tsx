import "../styles/globals.css";
import "@highlight-run/react/dist/highlight.css";
import type { AppProps } from "next/app";
import { H } from "highlight.run";
import Layout from "../components/Layout/Layout";
import { ErrorBoundary } from "@highlight-run/react";
import { HighlightOptions } from "../lib/highlight.config";
import { UserProvider } from "@supabase/supabase-auth-helpers/react";
import { supabaseClient } from "@supabase/supabase-auth-helpers/nextjs";
import {
  ColorScheme,
  ColorSchemeProvider,
  MantineProvider,
} from "@mantine/core";
import { useState } from "react";
import Head from "next/head";

H.init(process.env.NEXT_PUBLIC_HIGHLIGHT_PROJECT_ID, HighlightOptions);

function MyApp({ Component, pageProps }: AppProps) {
  const [colorScheme, setColorScheme] = useState<ColorScheme>("light");
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

  return (
    <>
      <Head>
        <title>Fig Issue</title>
      </Head>

      <ErrorBoundary>
        <ColorSchemeProvider
          colorScheme={colorScheme}
          toggleColorScheme={toggleColorScheme}
        >
          <MantineProvider
            theme={{ colorScheme }}
            withGlobalStyles
            withNormalizeCSS
          >
            <Layout>
              <UserProvider supabaseClient={supabaseClient}>
                <Component {...pageProps} />
              </UserProvider>
            </Layout>
          </MantineProvider>
        </ColorSchemeProvider>
      </ErrorBoundary>
    </>
  );
}

export default MyApp;
