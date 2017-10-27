<?php

trait AdminNewsTrait {
	// =======================
	// NEWS
	// =======================

	public function getNews($id = null) {

		if ($id) {
			if (Auth::user()->isDev()) {
				$news = Post::withTrashed()
					->findOrFail($id)
					->load("author");
			} else {
				$news = Post::findOrFail($id)
					->load("author");
			}
		} else {
			if (Auth::user()->isDev()) {
				$news = Post::with("author")
					->withTrashed()
					->orderBy("id", "desc")
					->paginate(15);
			} else {
				$news = Post::with("author")
					->orderBy("id", "desc")
					->paginate(15);
			}
			
		}

		$this->layout->content = View::make("admin.news")
			->with("news", $news)
			->with("id", $id);
	}

	public function postNews() {
		$title = Input::get("title");
		$text = Input::get("text");
		$header = Input::get("header");
		$user = Auth::user();
		$private = Input::get("private");

		if (!$user->canPostNews()) {
			Session::flash("status", "I can't let you do that.");
		} else {
			try {
				$post = Post::create([
					"user_id" => $user->id,
					"title" => $title,
					"text" => $text,
					"header" => $header,
					"private" => $private,
				]);
				

				$status = "News post added.";
			} catch (Exception $e) {
				$status = $e->getMessage();
			}
			
		}

		return Redirect::to("/admin/news")
			->with("status", $status);

	}

	public function putNews($id) {
		$title = Input::get("title");
		$text = Input::get("text");
		$header = Input::get("header");
		$private = Input::get("private");
		$user = Auth::user();

		if (!Auth::user()->canPostNews()) {
			Session::flash("status", "I can't let you do that.");
		} else {
			try {
				$post = Post::findOrFail($id);

				$post->title = $title;
				$post->header = $header;
				$post->text = $text;
				$post->private = $private;

				$post->save();

				$status = "News post $id updated.";
				
			} catch (Exception $e) {
				$status = $e->getMessage();
			}
			
		}

		return Redirect::to("/admin/news/$id")
			->with("status", $status);

	}

	public function deleteNews($id) {
		$user = Auth::user();
		if ($user->isAdmin()) {
			try {
				$post = Post::findOrFail($id);
				$title = $post->title;
				$post->delete();

				$status = "Post Deleted.";
				
			} catch (Exception $e) {
				$status = $e->getMessage();
			}
		}

		return Redirect::to("/admin/news")
			->with("status", $status);
	}
}
