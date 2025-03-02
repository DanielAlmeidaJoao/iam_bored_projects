package messages;

import dataClasses.LogEntry;
import pt.unl.fct.di.novasys.babel.generic.ProtoMessage;
import pt.unl.fct.di.novasys.network.data.Host;

import java.util.List;

public class AppendEntriesRPC extends ProtoMessage {
    public static final short ID = 200;
    private int term;
    private Host leaderId;
    private int previousLogIndex;
    private int previousLogTerm;
    private List<LogEntry> logEntriesToStore;
    private int leaderCommitIndex;

    public AppendEntriesRPC() {
        super(ID);
    }
}
