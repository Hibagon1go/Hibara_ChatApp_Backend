import type { NextPage } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import axiosJWTClient from "../../components/libs/requestAuth";
import { backend } from "../../constants/api";
import styles from "../../styles/Home.module.css";

const ChatRoom: NextPage = () => {
  const router = useRouter();
  const { room_id }: any = router.query;
  const FETCH_ALL_MSGS_API_URL = `${backend.BACKEND_BASE_URL}/api/msg?chat_room_id=${room_id}`;
  const MSG_API_URL = `${backend.BACKEND_BASE_URL}/api/msg`;
  const [text, setText] = useState("");
  const [editingText, setEditingText] = useState("");
  const [editingMsgID, setEditingMsgID] = useState("");
  const [willEdit, setWillEdit] = useState(false);
  interface ChatMsg {
    id: string;
    text: string;
    name: string;
    created_at: string;
    updated_at: string;
  }
  const [msgs, setMsgs] = useState<ChatMsg[]>([
    {
      id: "",
      text: "",
      name: "",
      created_at: "",
      updated_at: "",
    },
  ]);

  useEffect(() => {
    axiosJWTClient.get(FETCH_ALL_MSGS_API_URL).then((res) => {
      setMsgs(res.data);
      console.log(JSON.stringify(res.data));
    });
  }, []);

  const sendMsg = (text: string) => {
    axiosJWTClient
      .post(
        MSG_API_URL,
        {
          text: text,
          chat_room_id: room_id,
        }
      )
      .then((res) => {
        console.log(JSON.stringify(res.data));
      });
  };

  const editMsg = (msgID: string, text: string) => {
    axiosJWTClient
      .put(MSG_API_URL, {
        id: msgID,
        text: text,
      })
      .then((res) => {
        console.log(JSON.stringify(res.data));
      });
  };

  const deleteMsg = (msgID: string) => {
    axiosJWTClient
      .delete(MSG_API_URL, {
        data: { ID: msgID },
      })
      .then((res) => {
        console.log(JSON.stringify(res.data));
      });
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>チャット詳細</title>
        <meta name="description" content="Chat Detail" />
      </Head>

      <main className={styles.main}>
        {msgs.map((msg, key) => {
          return (
            <div key={key}>
              <h3>{msg.name}</h3>
              <p>
                {msg.text} {msg.created_at != msg.updated_at && "編集済み"}
              </p>
              <button
                onClick={() => {
                  setWillEdit(true);
                  setEditingMsgID(msg.id);
                  setEditingText(msg.text);
                }}
              >
                編集する
              </button>
              <button
                onClick={() => {
                  deleteMsg(msg.id);
                }}
              >
                削除する
              </button>
            </div>
          );
        })}
        {willEdit ? (
          <div>
            <input
              value={editingText}
              onChange={(event) => setEditingText(event.target.value)}
            />
            <button
              onClick={() => {
                editMsg(editingMsgID, editingText);
                setWillEdit(false);
              }}
            >
              メッセージを更新する！
            </button>
          </div>
        ) : (
          <div>
            <input
              value={text}
              onChange={(event) => setText(event.target.value)}
            />
            <button
              onClick={() => {
                sendMsg(text);
              }}
            >
              メッセージを送信する
            </button>
          </div>
        )}

        <button
          onClick={() => {
            axiosJWTClient
              .get(FETCH_ALL_MSGS_API_URL, {
              })
              .then((res) => {
                setMsgs(res.data);
                console.log(JSON.stringify(res.data));
              });
          }}
        >
          更新
        </button>
      </main>
    </div>
  );
};

export default ChatRoom;
