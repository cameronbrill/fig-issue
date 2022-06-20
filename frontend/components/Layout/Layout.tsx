import React, { ReactNode } from "react";
import { Header } from "./Header";

type LayoutProps = {
  children: ReactNode;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <div
        style={{
          position: "relative",
          minHeight: "91vh",
          overflowY: "scroll",
        }}
      >
        <Header
          links={[
            { link: "/", label: "Home" },
            { link: "/me", label: "Profile" },
          ]}
        />
        <main style={{ paddingBottom: "5vh", minHeight: "91vh" }}>
          {children}
        </main>
        <footer
          style={{
            height: "5vh",
            position: "absolute",
            bottom: "0",
            width: "100%",
          }}
        >
          Cameron and Nico
        </footer>
      </div>
    </>
  );
};

export default Layout;
