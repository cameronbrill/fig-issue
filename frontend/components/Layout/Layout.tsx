import React, { ReactNode } from "react";
import { Header } from "./Header";

import styles from "./Layout.module.css";

type LayoutProps = {
  children: ReactNode;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <div className={styles.container}>
        <Header
          links={[
            { link: "/", label: "Home" },
            { link: "/me", label: "Profile" },
          ]}
        />
        <main className={styles.main}>{children}</main>
        <footer className={styles.footer}>Cameron and Nico</footer>
      </div>
    </>
  );
};

export default Layout;
