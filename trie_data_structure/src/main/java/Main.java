import java.util.LinkedList;
import java.util.List;
import java.util.Scanner;

public class Main {
    public static String text = "A trie is a rooted tree used for storing associative arrays where keys are usually strings. Edges are often\n" +
            "labeled by individual symbols. Then common prefixes are factorized. Each node of the trie is associated\n" +
            "with a prefix of a string of the set of strings: concatenation of the labels of the path from the root to\n" +
            "the node. The root is associated with the empty string. Strings of the set are stored in terminal nodes\n" +
            "(leaves) but not in internal nodes. A trie can be seen as a Deterministic Finite Automaton.\n" +
            "Tries can be compacted. To get a compact trie from a trie, internal nodes with exactly one successor\n" +
            "are removed. Then labels of edges between remaining nodes are concatenated. Thus:\n" +
            "Edges are labeled by strings.\n" +
            "•Internal nodes have at least two children.\n" +
            "•Edges outgoing an internal node are labeled by strings starting with different symbols.";

    public static void main(String[] args) {
        TrieDST trieDST = new TrieDST();
        //String words [] = new String[] {"in", "integer", "interval", "string", "structure"};
        String words [] = text.split("\n");
        List<String> allWords = new LinkedList<>();
        for (String word : words) {
            trieDST.buildTrie(word.split(" "));
        }
        for (String word : allWords) {
            if(!trieDST.isInTrie(word)){
                System.out.println("NOT IN TRIE: "+word);
            }
        }

        Scanner scanner = new Scanner(System.in);
        String line = "";
        while(!"END".equals(line)){
            System.out.println("Type a new word to search: ");
            line = scanner.nextLine();
            System.out.println("Suggestions are: "+trieDST.getWords(line));
        }
    }
}