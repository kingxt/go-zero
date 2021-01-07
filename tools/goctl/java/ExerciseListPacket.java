package com.xhb.logic.http.packet.calcexercise;

import com.xhb.core.packet.HttpPacket;
import com.xhb.logic.http.packet.tracevisual.model.*;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.logic.http.packet.calcexercise.model.*;

public class ExerciseListPacket extends HttpPacket<ExerciseListResponse> {
	
	public ExerciseListPacket(ExerciseListRequest request) {
		super(request);
		this.request = request;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.GET;
    }

	@Override
    public String requestUri() {
        return "/calcexercise/exercise-list" + "?timestamp="request.getTimestamp();
    }
}
