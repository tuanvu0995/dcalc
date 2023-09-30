import { useAppState } from "../../context/AppStateContext";

export default function BottomBar() {
  const { state } = useAppState();

  const lineText = state.nodes.length > 1 ? "lines" : "line";

  const renderActionText = () => {
    if (state.clearing) return "[Enter to clear]";
    return "[Esc to clear]";
  }

  return (
    <div className="bottom-bar__container">
      <div>{renderActionText()}</div>
      <div>
        [{state.nodes.length}] {lineText}
      </div>
    </div>
  );
}
