import "../styles/globals.css";
import "@highlight-run/react/dist/highlight.css";
import type { AppProps } from "next/app";
import { H } from "highlight.run";
import { AuthProvider } from "../components/Auth/UserContext";
import { ErrorBoundary } from "@highlight-run/react";
import { HighlightOptions } from "../lib/highlight.config";

H.init(process.env.NEXT_PUBLIC_HIGHLIGHT_PROJECT_ID, HighlightOptions);

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ErrorBoundary>
      <AuthProvider>
        <Component {...pageProps} />
      </AuthProvider>
    </ErrorBoundary>
  );
}

export default MyApp;
