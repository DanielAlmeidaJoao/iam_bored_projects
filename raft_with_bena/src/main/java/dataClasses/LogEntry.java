package dataClasses;

public class LogEntry {
    private String command;
    private int term;

    public LogEntry(String command, int term) {
        this.command = command;
        this.term = term;
    }

    public String getCommand() {
        return command;
    }

    public int getTerm() {
        return term;
    }
}
