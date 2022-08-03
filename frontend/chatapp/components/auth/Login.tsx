import axios from "axios";
import type { NextPage } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { backend } from "../../constants/api";
import ColorfulButton from "../common/ColorfulButton";

const Login: NextPage = () => {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const LOGIN_URL = `${backend.BACKEND_BASE_URL}/login`;
  const ROOM_LIST_PATH = "/room_list";

  /**
  useEffect(() => {
    const isLogedIn = localStorage.getItem("token") !== null;
    if (isLogedIn) {
      router.push(ROOM_LIST_PATH);
    }
  }, []);
   */

  const loginHandler = () => {
    axios
      .post(LOGIN_URL, { email: email, password: password })
      .then((res) => {
        localStorage.setItem("token", res.data.token);
        router.push(ROOM_LIST_PATH);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="bg-white py-6 sm:py-8 lg:py-12">
      <div className="max-w-screen-2xl px-4 md:px-8 mx-auto">
        <h2 className="text-gray-800 text-2xl lg:text-3xl font-bold text-center mb-4 md:mb-8">
          Welcome back to ChatApp
        </h2>

        <form className="max-w-lg border rounded-lg mx-auto">
          <div className="flex flex-col gap-4 p-4 md:p-8">
            <div>
              <label
                htmlFor="email"
                className="inline-block text-gray-800 text-sm sm:text-base mb-2"
              >
                メールアドレス
              </label>
              <input
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                name="email"
                type="email"
                className="w-full bg-gray-50 text-gray-800 border focus:ring ring-indigo-300 rounded outline-none transition duration-100 px-3 py-2"
              />
            </div>

            <div>
              <label
                htmlFor="password"
                className="inline-block text-gray-800 text-sm sm:text-base mb-2"
              >
                パスワード
              </label>
              <input
                name="password"
                value={password}
                type="password"
                onChange={(e) => setPassword(e.target.value)}
                className="w-full bg-gray-50 text-gray-800 border focus:ring ring-indigo-300 rounded outline-none transition duration-100 px-3 py-2"
              />
            </div>
            <ColorfulButton onClick={loginHandler} buttonText="ログイン" />
          </div>

          <div className="flex justify-center items-center bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 p-4">
            <p className="text-white text-sm text-center">
              アカウントがない場合は{" "}
              <Link href="/signup">
                <a className="text-white text-sm">会員登録</a>
              </Link>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
