package org.raft;

import dataClasses.LogEntry;
import pt.unl.fct.di.novasys.babel.core.GenericProtocolExtension;
import pt.unl.fct.di.novasys.babel.exceptions.HandlerRegistrationException;
import pt.unl.fct.di.novasys.network.data.Host;

import java.io.IOException;
import java.util.List;
import java.util.Map;
import java.util.Properties;

public class PeerToPeerNode extends GenericProtocolExtension {
    public static final short PROTO_ID = 1;

    //Persistent state on all servers
    private int currentTerm;
    private String votedFor;
    private List<LogEntry> logEntries;

    //Volatile state on all servers:
    private int highestCommittedLogEntryIndex;
    private int highestAppliedLogEntryIndex;

    //Volatile state on leaders, reinitialized after election
    // index vs term ??

    private Map<Host,Integer> nextLogEntryIndexToSendToAllServers; //r (initialized to leaderlast log index + 1)
    private Map<Host,Integer> highestLogEntryIndexReplicatedToAllServer; //(initialized to 0, increases monotonically)

    public PeerToPeerNode(String protoName, short protoId) {
        super(protoName, protoId);
    }

    @Override
    public void init(Properties properties) throws HandlerRegistrationException, IOException {
        currentTerm = 0;
        highestAppliedLogEntryIndex = 0;
        highestCommittedLogEntryIndex = 0;
    }


}
