// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.calcexercise.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.packet.HttpData;
import com.xhb.core.response.HttpResponseData;

public class ExerciseDelRequest extends HttpData {
	private java.util.ArrayList<String> ids;

	@NotNull 
	public java.util.ArrayList<String> getIds() {
		return this.ids;
	}

	public void setIds(@NotNull java.util.ArrayList<String> ids) {
		this.ids = ids;
	}
}
