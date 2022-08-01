import NextDocument, { Head, Html, Main, NextScript } from "next/document";

type Props = {};

class Document extends NextDocument<Props> {
  render() {
    return (
      <Html lang="ja">
        <Head>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
        <footer className="flex flex-1 p-3 justify-center items-center  border-2">
          <a
            href="https://github.com/Hibagon1go"
            target="_blank"
            rel="noopener noreferrer"
          >
            Powered by Yoichiro Hibara
          </a>
        </footer>
      </Html>
    );
  }
}
export default Document;
