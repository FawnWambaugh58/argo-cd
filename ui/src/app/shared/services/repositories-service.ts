import {Observable} from 'rxjs';
import {models} from '../models';
import requests from './requests';

export class RepositoriesService {
    // ...
    public refs(repo: string, project?: string): Observable<models.RefsQueryResult> {
        return requests
            .get(`/repositories/${encodeURIComponent(repo)}/refs`)
            .query({project})
            .map(res => res.body as models.RefsQueryResult);
    }
    // ...
}