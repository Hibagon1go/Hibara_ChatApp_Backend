import { useRouter } from "next/router";

type RoomCardProps = {
  room: any;
  ROOM_PATH: string;
  leaveRoom?: (chatRoomID: string) => void;
  joinNewRoom?: (chatRoomID: string) => void;
  isNewRoom: boolean;
};

const RoomCard: React.FC<RoomCardProps> = (props) => {
  const { room, ROOM_PATH, isNewRoom } = props;
  const router = useRouter();
  return (
    <div className="m-2 bg-white">
      <div className="max-w-screen-2xl px-4 md:px-8 mx-auto">
        <div className="flex flex-col sm:flex-row justify-between items-center bg-gray-100 rounded-lg gap-4 p-4 md:p-8">
          <div>
            <h2 className="text-xl md:text-2xl font-bold">{room.name}</h2>
          </div>
          <div>
            {isNewRoom ? (
              <button
                onClick={() => {
                  props.joinNewRoom && props.joinNewRoom(room.chat_room_id);
                }}
                className="m-1 inline-block bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 focus-visible:ring ring-indigo-300 text-white text-sm md:text-base font-semibold text-center rounded-lg outline-none transition duration-100 px-8 py-3"
              >
                参加する
              </button>
            ) : (
              <>
                <button
                  onClick={() => {
                    router.push(ROOM_PATH + room.chat_room_id);
                  }}
                  className="m-1 inline-block bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 focus-visible:ring ring-indigo-300 text-white text-sm md:text-base font-semibold text-center rounded-lg outline-none transition duration-100 px-8 py-3"
                >
                  トークへ
                </button>
                <button
                  onClick={() => {
                    props.leaveRoom && props.leaveRoom(room.chat_room_id);
                  }}
                  className="m-1 inline-block bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700 focus-visible:ring ring-indigo-300 text-white text-sm md:text-base font-semibold text-center rounded-lg outline-none transition duration-100 px-8 py-3"
                >
                  退出する
                </button>
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default RoomCard;
