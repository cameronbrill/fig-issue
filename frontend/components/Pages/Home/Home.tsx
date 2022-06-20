import React from "react";
import { FeaturesSection } from "./FeaturesSection";
import { WelcomeHeroText } from "./WelcomHeroText";

import styles from "./Home.module.css";

export const HomePage: React.FC = () => {
  return (
    <div className={styles.container}>
      <WelcomeHeroText />
      <FeaturesSection />
    </div>
  );
};
