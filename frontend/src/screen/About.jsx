import { useEffect, useState } from "react";
import { Navigate } from "../../wailsjs/go/main/App";
import icon from "../assets/images/icon.png";

import { GetVersion } from "../../wailsjs/go/main/App";

export default function AboutScreen() {
  const [version, setVersion] = useState("---");

  useEffect(() => {
    GetVersion().then((v) => {
      setVersion(v);
    });
  }, []);

  return (
    <main className="container is-center" id="about">
      <figure class="image is-64x64">
        <img src={icon} />
      </figure>
      <h1 className="title is-3 my-2">DCalc</h1>
      <p className="mb-2">{version}</p>
      <p className="subtitle is-6 my-2">© 2023 TheDevLog</p>

      <div className="my-5">
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
