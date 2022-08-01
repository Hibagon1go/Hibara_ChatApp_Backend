type RoomCardProps = {
  onClick: () => void;
  buttonText: string;
};

const ColorfulButton: React.FC<RoomCardProps> = (props) => {
  const { onClick, buttonText } = props;
  return (
    <button
      type="button"
      onClick={onClick}
      className="text-white rounded p-4 bg-gradient-to-r from-green-400 to-blue-500 hover:from-pink-500 hover:to-yellow-500 ..."
    >
      {buttonText}
    </button>
  );
};

export default ColorfulButton;
