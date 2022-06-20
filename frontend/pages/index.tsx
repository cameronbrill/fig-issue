import type { NextPage } from "next";
import { HomePage } from "../components/Pages/Home/Home";
import styles from "../styles/Home.module.css";

const Home: NextPage = () => {
  return (
    <section className={styles.main}>
      <HomePage />
    </section>
  );
};

export default Home;
