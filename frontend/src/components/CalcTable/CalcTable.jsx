import { useEffect, useRef } from "react";
import { useAppState } from "../../context/AppStateContext";
import LineItem from "./LineItem";
import { OnTyping } from "../../../wailsjs/go/main/App";
import SummaryLine from "./SummaryLine";

export default function CalcTable() {
  const { state } = useAppState();
  const bottomRef = useRef(null);

  const handleOnKeyPress = (e) => {
    e.preventDefault();
    console.log("handleOnKeyPress", e.key);
    OnTyping(e.key);
  };

  useEffect(() => {
    document.addEventListener("keypress", handleOnKeyPress);
    return () => document.removeEventListener("keypress", handleOnKeyPress);
  }, []);

  useEffect(() => {
    if (bottomRef.current) {
      console.log("scrollIntoView");
      bottomRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [state]);

  return (
    <div className="calc-table__container">
      <div className="table__container">
        <table className="table">
          <tbody>
            {state.nodes.map((node, index) => (
              <LineItem key={node.name} index={index} node={node} />
            ))}
          </tbody>
        </table>
        <SummaryLine />
        <div ref={bottomRef}></div>
      </div>
    </div>
  );
}
