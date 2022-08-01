import Image from "next/image";
import Link from "next/link";

const NotFoundPage = () => {
  return (
    <div className="bg-white py-6 sm:py-8 lg:py-12">
      <div className="max-w-screen-2xl px-4 md:px-8 mx-auto">
        <div className="w-full sm:w-96 h-96 flex justify-center items-center bg-gray-100 shadow-lg rounded-lg overflow-hidden relative mx-auto">
          <Image
            src="/404.png"
            loading="lazy"
            alt="Photo by Connor Botts"
            layout='fill'
            className="w-full h-full object-cover object-center absolute inset-0"
          />

          <div className="flex flex-col justify-center items-center relative p-8 md:p-16">
            <h1 className="text-black text-2xl md:text-3xl lg:text-4xl font-bold text-center mb-2">
              404
            </h1>

            <p className="text-gray-600 md:text-lg text-center mb-8">
              お探しのページは見つかりません.
            </p>

            <Link href="/">
              <a className="inline-block bg-gray-200 hover:bg-gray-300 focus-visible:ring ring-indigo-300 text-white active:text-gray-700 text-sm md:text-base font-semibold text-center rounded-lg outline-none transition duration-100 px-8 py-3">
                トップに戻る
              </a>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NotFoundPage;
