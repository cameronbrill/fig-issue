import React, { ReactNode } from "react";
import { Footer } from "./Footer";
import { Header } from "./Header";

import styles from "./Layout.module.css";

type LayoutProps = {
  children: ReactNode;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const links = [
    { link: "/", label: "Home" },
    { link: "/me", label: "Profile" },
  ];
  return (
    <>
      <Header links={links} />
      <div className={styles.container}>
        <main className={styles.main}>{children}</main>
      </div>
      <Footer links={links} />
    </>
  );
};

export default Layout;
