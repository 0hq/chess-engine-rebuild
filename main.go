package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/0hq/chess"
)

/*

// Replace position with board.
// Evaluation function.
// Add in auto-testing using EPD files.
// Test different sorting algorithms.
// Change checkmate value to not be max int.
// Engine now prioritize the shortest checkmate.
// Add in Eigenmann rapid engine test.
// Piece square tables.
// Turn iterative deepening into an engine.
// Change engine struct to be smarter.
// Make iterative deepening play in time.
// Clone go chess package to customize it.
// Killer moves.
// Mobility is done, but seems to be a massive perf. hit.
// Pick and sort, changes engine structure.
// Hash MVV/LVA
// Transposition tables.
// TT Move ordering.
// Draw detection

MTD(f)
Opening book analysis. (fuck)
Lines via linked list.
PSTO evaluation.
	Tapered eval
	Endgames
Better move ordering.
   SEE
   History heuristic
Parallel search via LAZY SMP

Investigate 0.3.1 beating 0.3.4 in performance. Probably just hashing slowdowns from static?
Make total nodes trustable.

UCI compatibility. Ugh, this sucks. I might give up on this and do a web server.

Null Move Pruning.
Check extentions.
Investigate mobility.
Passed pawns, blocked pawns, etc.
Investigate transposition work saving.

Extensive testing positions.
	Better test parsing.
	Zugzwang.

Magic bitboards.
Why isn't the parallel version faster for perft?
Aspirations?

*/

func init() {
	out("Initializing engine...")
	InitZobrist()
	// create new log file that doesn't exist
	if !production_mode {
		for i := 0; ; i++ {
			// create file name from timestamp date and hour
			date := time.Now().Format("2006-01-02")
			filename := fmt.Sprintf("logs/%s-%d.log", date, i)
			_, err := os.Stat(filename)
			if os.IsNotExist(err) {
				f, err := os.Create(filename)
				if err != nil {
					log.Fatal(err)
				}
				log.SetOutput(f)
				break
			}
		}
		out("File initialization.")
	}

	out("Version", runtime.Version())
    out("NumCPU", runtime.NumCPU())
    out("GOMAXPROCS", runtime.GOMAXPROCS(0))
	out("Initialization complete.")
	out()
	if production_mode {
		out("Production mode.")
	} else {
		out("WARNING: Development mode.")
	}
	if QUIET_MODE {
		out("WARNING: Quiet mode!")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	out("Running main program.", "\n")
	defer exit()

	// engine := new_engine(&engine_0dot3dot4, 1, nil)
	// simple_tests(engine)
	// eigenmann_tests(engine)

	// mini_performance_challenge()
	// mini_test_transposition()
	// mini_self_challenge()
	// mini_challenge_stockfish()
}

func exit()	{
	out("Exiting engine.")
}

func mini_test_transposition() {
	game := game_from_fen("6r1/pppk4/3p4/8/2PnPp1Q/7P/PP4r1/R5RK b - - 1 24")
	engine2 := wrap_engine(&engine_0dot3, 5, game)
	// stop_at_depth = 7
	engine2.Run_Engine(game.Position())
	// do_tt_output = true
	// stop_at_depth = 12
	engine2.Run_Engine(game.Position())
}

func mini_performance_challenge() {
	out("Performance challenge!", "\n")
	game := game_from_fen("r3kb1r/2p2ppp/p1p5/4P3/4n3/8/PPP3PP/RNB2RK1 b kq - 0 13")
	engine1 := wrap_engine(&engine_0dot3dot4, 10, game)
	engine2 := wrap_engine(&engine_0dot3dot1, 10, game)
	engine1.Run_Engine(game.Position())
	engine1.Run_Engine(game.Position())

	engine2.Run_Engine(game.Position())
	engine2.Run_Engine(game.Position())

}

func mini_self_challenge() {
	game := game_from_fen(CHESS_START_POSITION)
	engine1 := wrap_engine(&engine_0dot3dot4, 5, game)
	engine2 := wrap_engine(&engine_0dot3dot4, 5, game)
	challenge_self(engine1, engine2, game)
}

func mini_challenge_manual() {
	game := game_from_fen(CHESS_START_POSITION)
	engine := wrap_engine(&engine_0dot2, 15, game)
	challenge_manual(engine, chess.Black, game)
}

func mini_challenge_stockfish() {
	game := game_from_fen(CHESS_START_POSITION)
	engine := wrap_engine(&engine_0dot3dot2, 15, game)
	challenge_stockfish(engine, chess.White, game)
}

func mini_simple() {
	engine := new_engine(&engine_0dot2dot1, 15, nil)
	simple_tests(engine)
}

// current will vs antikythera 2kr1b1r/ppq2ppp/2n1bn2/1R2p1B1/8/2NP1N2/P1P2PPP/3QR1K1 w - - 10 15
// out(evaluate_position_v3(game_from_fen("rnb1k2r/ppp3pp/8/5p2/3qnP2/2N5/PPPNQbPP/R1BK1B1R w kq - 18 19").Position(), 0, 0, 1))
// test(engine, "6r1/pppk4/3p4/8/2PnPp1Q/7P/PP4r1/R5RK b - - 1 24", "g2g1", false)