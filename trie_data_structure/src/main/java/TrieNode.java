import java.util.*;

public class TrieNode {
    private TrieNode parent;
    private char letter;
    private Map<Character,TrieNode> children;
    private Set<String> words;
    private String exactWord;
    public TrieNode(char letter){
        children = new HashMap<>();
        words = new HashSet<>();
        this.letter = letter;
    }

    public boolean isLeaf(){
        return children.isEmpty();
    }
    public boolean isInternalNode(){
        return !isLeaf();
    }
    public TrieNode getChild(char key){
        return children.get(key);
    }
    public TrieNode addChild(char key, TrieNode node){
        return children.put(key,node);
    }
    public boolean hasChild(char child){
        return children.containsKey(child);
    }
    public void setWords(String value){
        words.add(value);
    }
    public String getWords(){
        return words.toString();
    }
}
