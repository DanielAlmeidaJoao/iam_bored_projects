public class TrieDST {
    private TrieNode root;
    public TrieDST(){
        root = new TrieNode(' ');
    }
    public void buildTrie(String [] words){
        TrieNode temp, auxParent;
        for (String word : words) {
            auxParent = root;
            for (int i = 0; i < word.length(); i++) {
                char c = word.charAt(i);
                temp = auxParent.getChild(c);
                if(temp == null){
                    temp = new TrieNode(c);
                    auxParent.addChild(c, temp );
                }
                temp.setWords(word);
                auxParent = temp;
            }
            //auxParent.setWords(word);
        }
    }

    private TrieNode getRoot(){
        return root;
    }

    public boolean isInTrie(String word){
        return isInTrie(root,word)!=null;
    }

    private TrieNode isInTrie(TrieNode parent,String word){
        if (word == null){
            return null;
        } else if (word.isEmpty()){
            return parent;
        } else {
            TrieNode trieNode = parent.getChild(word.charAt(0));
            if (trieNode == null){
                return null;
            } else return isInTrie(trieNode,word.substring(1));
        }
    }

    public String getWords(String word){
        TrieNode node = isInTrie(root,word);
        if (node == null){
            return null;
        } else {
            return node.getWords();
        }
    }


}
