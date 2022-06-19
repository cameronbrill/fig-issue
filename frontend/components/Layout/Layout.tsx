import Head from "next/head";
import Link from "next/link";
import React, { ReactNode } from "react";

type LayoutProps = {
  children: ReactNode;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div>
      <Head>
        <title>Fig Issue</title>
      </Head>
      <header
        style={{
          display: "flex",
          marginRight: "5vw",
          marginLeft: "5vw",
          marginTop: "1vh",
          height: "5vh",
          position: "fixed",
        }}
      >
        <h1
          style={{
            minWidth: "fit-content",
            padding: "10px",
          }}
        >
          Fig Issue
        </h1>
        <div
          style={{
            display: "flex",
            width: "100%",
            justifyContent: "left",
            alignItems: "center",
          }}
        >
          <Link href="/">home</Link>
          <Link href="/me">me</Link>
        </div>
      </header>
      <div
        style={{
          overflowY: "scroll",
        }}
      >
        <main>{children}</main>
        <footer
          style={{
            height: "5vh",
          }}
        >
          Cameron and Nico
        </footer>
      </div>
    </div>
  );
};

export default Layout;
