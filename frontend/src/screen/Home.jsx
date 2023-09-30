import BottomBar from "../components/BottomBar";
import CalcTable from "../components/CalcTable";

export default function HomeScreen() {

  return (
    <main className="container" id="home">
      <CalcTable />
      <BottomBar />
    </main>
  );
}
