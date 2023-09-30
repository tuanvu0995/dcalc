import { Navigate } from "../../wailsjs/go/main/App";

export default function GuideScreen() {
  return (
    <main className="container p-2" id="about">
      <h1 className="title is-5 my-2">Guide</h1>

      <div className="scrollable">
        <div className="content">
          <p>
            Type the calculation and press <strong>[Enter]</strong>. The result
            is displayed in the cell.
          </p>
          <p>
            Press <strong>[Esc]</strong> to jump to clear mode. Press{" "}
            <strong>[Enter]</strong> to clear or press <strong>[Esc]</strong>
            to cancel.
          </p>

          <p>Press <strong>[Cmd+E]</strong> or <strong>[Ctrl+E]</strong> to export csv file.</p>

          <h2>Operators</h2>
          <p>Operators are the symbols that perform the calculation.</p>
          <p>Operators are: +, -, *, /</p>

          <h2>Functions</h2>
          <p>Functions are the symbols that perform the calculation.</p>
          <p>
            Functions are: abs, ceil, cos, tan, log, max, min, pow, sqrt.
          </p>
        </div>
      </div>

      <div className="is-flex is-justify-content-center py-4">
        <button
          className="button is-small is-primary is-outlined"
          onClick={() => Navigate("/")}
        >
          Back
        </button>
      </div>
    </main>
  );
}
