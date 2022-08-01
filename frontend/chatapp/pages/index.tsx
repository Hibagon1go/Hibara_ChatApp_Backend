import type { NextPage } from "next";
import Head from "next/head";
import Login from "../components/auth/Login";
import styles from "../styles/Home.module.css";

const SignUpPage: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>ログイン</title>
        <meta
          name="description"
          content="Login Page"
        />
        
      </Head>

      <main className={styles.main}>
        <Login />
      </main>
    </div>
  );
};

export default SignUpPage;
