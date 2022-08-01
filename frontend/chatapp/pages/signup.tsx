import type { NextPage } from "next";
import Head from "next/head";
import SignUp from "../components/auth/SignUp";
import styles from "../styles/Home.module.css";

const SignUpPage: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>会員登録</title>
        <meta
          name="description"
          content="SignUp Page"
        />
      </Head>

      <main className={styles.main}>
        <SignUp />
      </main>
    </div>
  );
};

export default SignUpPage;
