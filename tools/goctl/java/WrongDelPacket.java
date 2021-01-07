package com.xhb.logic.http.packet.calcexercise;

import com.xhb.core.packet.HttpPacket;
import com.xhb.logic.http.packet.tracevisual.model.*;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.calcexercise.model.*;

public class WrongDelPacket extends HttpPacket<WrongDelResponse> {
	
	public WrongDelPacket(WrongDelRequest request) {
		super(request);
		this.request = request;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.DELETE;
    }

	@Override
    public String requestUri() {
        return "/calcexercise/wrong-del";
    }
}
